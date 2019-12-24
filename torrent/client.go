package torrent

import (
    "encoding/json"
    format "fmt"
    "github.com/plexmediamanager/micro-torrent/errors"
    "github.com/plexmediamanager/service/helpers"
    "golang.org/x/net/publicsuffix"
    "io/ioutil"
    "net/http"
    "net/http/cookiejar"
    "net/url"
    "strings"
)

type Client struct {
    Authenticated       bool

    host                string
    username            string
    password            string
    client              *http.Client
    cookieJar           http.CookieJar
}

// Initialize torrent client configuration
func Initialize() *Client {
    client := &Client{
        host:       helpers.GetEnvironmentVariableAsString("TORRENT_HOST", "http://localhost"),
        username:   helpers.GetEnvironmentVariableAsString("TORRENT_LOGIN", ""),
        password:   helpers.GetEnvironmentVariableAsString("TORRENT_PASSWORD", ""),
    }
    client.cookieJar, _ = cookiejar.New(&cookiejar.Options{
        PublicSuffixList: publicsuffix.List,
    })
    client.client = &http.Client{
        Jar: client.cookieJar,
    }
    return client
}

// Perform authentication procedure
func (client *Client) Authenticate() error {
    credentials := make(map[string]string)
    credentials["username"] = client.username
    credentials["password"] = client.password

    response, err := client.sendPostRequest(ENDPOINT_AUTHENTICATION_LOGIN, credentials)
    if err != nil {
        return errors.TorrentAuthenticationError.ToError(err)
    } else if response.StatusCode != 200 {
        return errors.TorrentAuthenticationError.ToError(err)
    }

    client.Authenticated = true
    if cookies := response.Cookies(); len(cookies) > 0 {
        cookieUrl, _ := url.Parse("http://localhost")
        client.cookieJar.SetCookies(cookieUrl, cookies)
    }
    return nil
}


// Build request URL
func (client *Client) buildRequestUrl(url string) string {
    return format.Sprintf(url, client.host)
}

func (client *Client) sendGetRequest(endpoint string, options map[string]string) (*http.Response, error) {
    return client.sendRequest("GET", endpoint, options)
}

// Send POST request to API
func (client *Client) sendPostRequest(endpoint string, options map[string]string) (*http.Response, error) {
    return client.sendRequest("POST", endpoint, options)
}

// Send request
func (client *Client) sendRequest(method string, endpoint string, options map[string]string) (*http.Response, error) {
    var request *http.Request
    var err error
    if options != nil {
        requestForm := url.Values{}
        for key, value := range options {
            requestForm.Add(key, value)
        }
        if method == "POST" {
            request, err = http.NewRequest(method, client.buildRequestUrl(endpoint), strings.NewReader(requestForm.Encode()))
        } else {
            request, err = http.NewRequest(method, client.buildRequestUrl(endpoint) + "?" + requestForm.Encode(), nil)
        }
    } else {
        request, err = http.NewRequest(method, client.buildRequestUrl(endpoint), nil)
    }

    if err != nil {
        return nil, errors.RequestCreationError.ToError(err)
    }

    if method == "POST" {
        request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    }
    request.Header.Set("User-Agent", "FreedomCore Micro Torrent")


    response, err := client.client.Do(request)
    if err != nil {
        return nil, errors.RequestCouldNotBeCompleted.ToErrorWithArguments(err, client.buildRequestUrl(endpoint))
    }
    return response, nil
}

// Convert response to bytes
func (client *Client) responseToBytes(request *http.Response) []byte {
    response := make([]byte, 256)
    response, _ = ioutil.ReadAll(request.Body)
    return response
}

func (client *Client) get(unmarshal interface{}, endpoint string, options map[string]string) (interface {}, error) {
    response, err := client.sendGetRequest(endpoint, options)
    if err != nil {
        return nil, err
    }
    bytes := client.responseToBytes(response)
    err = json.Unmarshal(bytes, &unmarshal)
    if err != nil {
        format.Println(err)
        return nil, errors.UnmarshalError.ToError(err)
    }
    return unmarshal, nil
}