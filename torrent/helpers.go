package torrent

import (
    format "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "net/http/httputil"
)

func printResponse(body io.ReadCloser) {
    response := make([]byte, 256)
    response, _ = ioutil.ReadAll(body)
    format.Println("response:", string(response))
}

func printRequest(request *http.Request) error {
    dump, err := httputil.DumpRequest(request, true)
    if err != nil {
        return err
    }
    format.Println("request:", string(dump))
    return nil
}