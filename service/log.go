package service

import (
    "context"
    microClient "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-torrent/proto"
    "github.com/plexmediamanager/micro-torrent/torrent"
)

func LogEntries (client microClient.Client, options map[string]string) (*torrent.LogEntry, error) {
    var result *torrent.LogEntry
    service := GetTorrentService(client)
    parameters := &proto.TorrentOptions { Options: options }
    response, err := service.LogEntries(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}
