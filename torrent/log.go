package torrent

type LogEntry []struct {
    ID                                  uint64              `json:"id"`
    Message                             string              `json:"message"`
    Timestamp                           uint64              `json:"timestamp"`
    Type                                uint64              `json:"type"`
}

// Get log
func (client *Client) LogEntries(options map[string]string) (*LogEntry, error) {
    var response LogEntry
    newOptions := map[string]string {
        "normal":           "true",
        "info":             "true",
        "warning":          "true",
        "critical":         "true",
        "last_known_id":    "-1",
    }
    for key, value := range options {
        newOptions[key] = value
    }
    result, err := client.get(&response, ENDPOINT_LOG_GET, newOptions)
    return result.(*LogEntry), err
}
