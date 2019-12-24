package service

import (
    "context"
    microClient "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-torrent/proto"
    "github.com/plexmediamanager/micro-torrent/torrent"
)

func ApplicationVersion (client microClient.Client) (*torrent.ApplicationVersion, error) {
    var result *torrent.ApplicationVersion
    service := GetTorrentService(client)
    parameters := &proto.TorrentEmpty {}
    response, err := service.ApplicationVersion(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func ApplicationAPIVersion (client microClient.Client) (*torrent.ApplicationAPIVersion, error) {
    var result *torrent.ApplicationAPIVersion
    service := GetTorrentService(client)
    parameters := &proto.TorrentEmpty {}
    response, err := service.ApplicationAPIVersion(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func ApplicationBuildInformation (client microClient.Client) (*torrent.ApplicationBuildInfo, error) {
    var result *torrent.ApplicationBuildInfo
    service := GetTorrentService(client)
    parameters := &proto.TorrentEmpty {}
    response, err := service.ApplicationBuildInformation(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func ApplicationPreferences (client microClient.Client) (*torrent.ApplicationPreferences, error) {
    var result *torrent.ApplicationPreferences
    service := GetTorrentService(client)
    parameters := &proto.TorrentEmpty {}
    response, err := service.ApplicationPreferences(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}
