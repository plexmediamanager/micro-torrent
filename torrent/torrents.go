package torrent

import (
    format "fmt"
    "strings"
)

type Torrent []struct {
    AddedOn                 int64       `json:"added_on"`
    AmountLeft              int64       `json:"amount_left"`
    AutoTMM                 bool        `json:"auto_tmm"`
    Availability            float64     `json:"availability"`
    Category                string      `json:"category"`
    Completed               int64       `json:"completed"`
    CompletionOn            int64       `json:"completion_on"`
    DownloadLimit           int64       `json:"dl_limit"`
    DownloadSpeed           int64       `json:"dlspeed"`
    Downloaded              int64       `json:"downloaded"`
    DownloadedSession       int64       `json:"downloaded_session"`
    EstimatedTimeLeft       int64       `json:"eta"`
    FLPiecePrio             bool        `json:"f_l_piece_prio"`
    ForceStart              bool        `json:"force_start"`
    Hash                    string      `json:"hash"`
    LastActivity            int64       `json:"last_activity"`
    MagnetURI               string      `json:"magnet_uri"`
    MaxRatio                int64       `json:"max_ratio"`
    MaxSeedingTime          int64       `json:"max_seeding_time"`
    Name                    string      `json:"name"`
    NumComplete             int64       `json:"num_complete"`
    NumIncomplete           int64       `json:"num_incomplete"`
    NumberOfLeechers        int64       `json:"num_leechs"`
    NumberOfSeeders         int64       `json:"num_seeds"`
    Priority                int64       `json:"priority"`
    Progress                float64     `json:"progress"`
    Ratio                   int64       `json:"ratio"`
    RatioLimit              int64       `json:"ratio_limit"`
    SavePath                string      `json:"save_path"`
    SeedingTimeLimit        int64       `json:"seeding_time_limit"`
    SeenComplete            int64       `json:"seen_complete"`
    SequentialDownload      bool        `json:"seq_dl"`
    Size                    int64       `json:"size"`
    State                   string      `json:"state"`
    SuperSeeding            bool        `json:"super_seeding"`
    Tags                    string      `json:"tags"`
    TimeActive              int64       `json:"time_active"`
    TotalSize               int64       `json:"total_size"`
    Tracker                 string      `json:"tracker"`
    UploadLimit             int64       `json:"up_limit"`
    Uploaded                int64       `json:"uploaded"`
    UploadedSession         int64       `json:"uploaded_session"`
    UploadSpeed             int64       `json:"upspeed"`
}

type TorrentProperties struct {
    AdditionDate                    int64  `json:"addition_date"`
    Comment                         string `json:"comment"`
    CompletionDate                  int64  `json:"completion_date"`
    CreatedBy                       string `json:"created_by"`
    CreationDate                    int64  `json:"creation_date"`
    DownloadLimit                   int64  `json:"dl_limit"`
    DownloadSpeed                   int64  `json:"dl_speed"`
    DownloadSpeedAverage            int64  `json:"dl_speed_avg"`
    EstimatedTimeLeft               int64  `json:"eta"`
    LastSeen                        int64  `json:"last_seen"`
    ConnectionsCount                int64  `json:"nb_connections"`
    ConnectionsLimit                int64  `json:"nb_connections_limit"`
    Peers                           int64  `json:"peers"`
    PeersTotal                      int64  `json:"peers_total"`
    PieceSize                       int64  `json:"piece_size"`
    PiecesHave                      int64  `json:"pieces_have"`
    PiecesNumber                    int64  `json:"pieces_num"`
    Reannounce                      int64  `json:"reannounce"`
    SavePath                        string `json:"save_path"`
    SeedingTime                     int64  `json:"seeding_time"`
    Seeds                           int64  `json:"seeds"`
    SeedsTotal                      int64  `json:"seeds_total"`
    ShareRatio                      int64  `json:"share_ratio"`
    TimeElapsed                     int64  `json:"time_elapsed"`
    TotalDownloaded                 int64  `json:"total_downloaded"`
    TotalDownloadedSession          int64  `json:"total_downloaded_session"`
    TotalSize                       int64  `json:"total_size"`
    TotalUploaded                   int64  `json:"total_uploaded"`
    TotalUploadedSession            int64  `json:"total_uploaded_session"`
    TotalWasted                     int64  `json:"total_wasted"`
    UploadLimit                     int64  `json:"up_limit"`
    UploadSpeed                     int64  `json:"up_speed"`
    UploadSpeedAverage              int64  `json:"up_speed_avg"`
}

type TorrentTrackers []struct {
    Message                 string      `json:"msg"`
    NumDownloaded           int64       `json:"num_downloaded"`
    LeechesCount            int64       `json:"num_leeches"`
    PeersCount              int64       `json:"num_peers"`
    SeedersCount            int64       `json:"num_seeds"`
    Status                  int64       `json:"status"`
    //Tier                    int64       `json:"tier"` // TODO: so they can give back either INT or STRING, WTF?
    URL                     string      `json:"url"`
}

type TorrentContents []struct {
    Availability            int64       `json:"availability"`
    IsSeed                  bool        `json:"is_seed"`
    Name                    string      `json:"name"`
    PieceRange              []int64     `json:"piece_range"`
    Priority                int64       `json:"priority"`
    Progress                float64     `json:"progress"`
    Size                    int64       `json:"size"`
}

type TorrentPiecesState []int64

type TorrentPiecesHashes []string

// Get list of active torrents
func (client *Client) TorrentList() (*Torrent, error) {
    var response Torrent
    result, err := client.get(&response, ENDPOINT_TORRENT_LIST, nil)
    return result.(*Torrent), err
}

// Get properties of a specific torrent
func (client *Client) TorrentProperties(hash string) (*TorrentProperties, error) {
    var response TorrentProperties
    result, err := client.get(&response, ENDPOINT_TORRENT_PROPERTIES, map[string]string {
        "hash":     hash,
    })
    return result.(*TorrentProperties), err
}

// Get trackers of a specific torrent
func (client *Client) TorrentTrackers(hash string) (*TorrentTrackers, error) {
    var response TorrentTrackers
    result, err := client.get(&response, ENDPOINT_TORRENT_TRACKERS, map[string]string {
        "hash":     hash,
    })
    return result.(*TorrentTrackers), err
}

// Get contents of a specific torrent
func (client *Client) TorrentContents(hash string) (*TorrentContents, error) {
    var response TorrentContents
    result, err := client.get(&response, ENDPOINT_TORRENT_CONTENTS, map[string]string {
        "hash":     hash,
    })
    return result.(*TorrentContents), err
}

// Get pieces states of a specific torrent
func (client *Client) TorrentPiecesState(hash string) (*TorrentPiecesState, error) {
    var response TorrentPiecesState
    result, err := client.get(&response, ENDPOINT_TORRENT_PIECES_STATES, map[string]string {
        "hash":     hash,
    })
    return result.(*TorrentPiecesState), err
}

// Get pieces hashes of a specific torrent
func (client *Client) TorrentPiecesHashes(hash string) (*TorrentPiecesHashes, error) {
    var response TorrentPiecesHashes
    result, err := client.get(&response, ENDPOINT_TORRENT_PIECES_HASHES, map[string]string {
        "hash":     hash,
    })
    return result.(*TorrentPiecesHashes), err
}

// Pause specified torrents
func (client *Client) TorrentPause(hashes ...string) error {
    _, err := client.sendPostRequest(ENDPOINT_TORRENT_PAUSE, map[string]string {
        "hashes":       strings.Join(hashes, "|"),
    })
    return err
}

// Resume specified torrents
func (client *Client) TorrentResume(hashes ...string) error {
    _, err := client.sendPostRequest(ENDPOINT_TORRENT_RESUME, map[string]string {
        "hashes":       strings.Join(hashes, "|"),
    })
    return err
}

// Delete specified torrents
func (client *Client) TorrentDelete(hashes ...string) error {
    _, err := client.sendPostRequest(ENDPOINT_TORRENT_DELETE, map[string]string {
        "hashes":       strings.Join(hashes, "|"),
    })
    return err
}

// Delete specified torrents and all downloaded files
func (client *Client) TorrentDeleteWithFiles(hashes ...string) error {
    _, err := client.sendPostRequest(ENDPOINT_TORRENT_DELETE, map[string]string {
        "hashes":       strings.Join(hashes, "|"),
        "deleteFiles":  "true",
    })
    return err
}

// Re-Check specified torrents
func (client *Client) TorrentReCheck(hashes ...string) error {
    _, err := client.sendPostRequest(ENDPOINT_TORRENT_RECHECK, map[string]string {
        "hashes":       strings.Join(hashes, "|"),
    })
    return err
}

// Re-Announce specified torrents
func (client *Client) TorrentReAnnounce(hashes ...string) error {
    _, err := client.sendPostRequest(ENDPOINT_TORRENT_REANNOUNCE, map[string]string {
        "hashes":       strings.Join(hashes, "|"),
    })
    return err
}

// Increase priority for specified torrents
func (client *Client) TorrentIncreasePriority(hashes ...string) error {
    _, err := client.sendPostRequest(ENDPOINT_TORRENT_INCREASE_PRIORITY, map[string]string {
        "hashes":       strings.Join(hashes, "|"),
    })
    return err
}

// Decrease priority for specified torrents
func (client *Client) TorrentDecreasePriority(hashes ...string) error {
    _, err := client.sendPostRequest(ENDPOINT_TORRENT_DECREASE_PRIORITY, map[string]string {
        "hashes":       strings.Join(hashes, "|"),
    })
    return err
}

// Set highest priority for specified torrents
func (client *Client) TorrentSetHighestPriority(hashes ...string) error {
    _, err := client.sendPostRequest(ENDPOINT_TORRENT_MAX_PRIORITY, map[string]string {
        "hashes":       strings.Join(hashes, "|"),
    })
    return err
}

// Set lowest priority for specified torrents
func (client *Client) TorrentSetLowestPriority(hashes ...string) error {
    _, err := client.sendPostRequest(ENDPOINT_TORRENT_MIN_PRIORITY, map[string]string {
        "hashes":       strings.Join(hashes, "|"),
    })
    return err
}

// Set lowest priority for specified torrents
func (client *Client) TorrentSetFilePriority(hash string, priority uint64, files ...string) error {
    _, err := client.sendPostRequest(ENDPOINT_TORRENT_MIN_PRIORITY, map[string]string {
        "hash":     hash,
        "priority": format.Sprint(priority),
        "id":       strings.Join(files, "|"),
    })
    return err
}