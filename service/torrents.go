package service

import (
    "context"
    microClient "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-torrent/proto"
    "github.com/plexmediamanager/micro-torrent/torrent"
)

func TorrentList (client microClient.Client) (*torrent.Torrent, error) {
    var result *torrent.Torrent
    service := GetTorrentService(client)
    parameters := &proto.TorrentEmpty {}
    response, err := service.TorrentList(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TorrentProperties(client microClient.Client, hash string) (*torrent.TorrentProperties, error) {
    var result *torrent.TorrentProperties
    service := GetTorrentService(client)
    parameters := &proto.TorrentString { Value: hash }
    response, err := service.TorrentProperties(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TorrentTrackers(client microClient.Client, hash string) (*torrent.TorrentTrackers, error) {
    var result *torrent.TorrentTrackers
    service := GetTorrentService(client)
    parameters := &proto.TorrentString { Value: hash }
    response, err := service.TorrentTrackers(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TorrentContents(client microClient.Client, hash string) (*torrent.TorrentContents, error) {
    var result *torrent.TorrentContents
    service := GetTorrentService(client)
    parameters := &proto.TorrentString { Value: hash }
    response, err := service.TorrentContents(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TorrentPiecesState(client microClient.Client, hash string) (*torrent.TorrentPiecesState, error) {
    var result *torrent.TorrentPiecesState
    service := GetTorrentService(client)
    parameters := &proto.TorrentString { Value: hash }
    response, err := service.TorrentPiecesState(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TorrentPiecesHashes(client microClient.Client, hash string) (*torrent.TorrentPiecesHashes, error) {
    var result *torrent.TorrentPiecesHashes
    service := GetTorrentService(client)
    parameters := &proto.TorrentString { Value: hash }
    response, err := service.TorrentPiecesHashes(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func TorrentPause(client microClient.Client, hashes ...string) error {
    service := GetTorrentService(client)
    parameters := &proto.TorrentStrings{ Values: hashes}
    _, err := service.TorrentPause(context.TODO(), parameters)
    return err
}

func TorrentResume(client microClient.Client, hashes ...string) error {
    service := GetTorrentService(client)
    parameters := &proto.TorrentStrings{ Values: hashes}
    _, err := service.TorrentResume(context.TODO(), parameters)
    return err
}

func TorrentDelete(client microClient.Client, hashes ...string) error {
    service := GetTorrentService(client)
    parameters := &proto.TorrentStrings{ Values: hashes}
    _, err := service.TorrentDelete(context.TODO(), parameters)
    return err
}

func TorrentDeleteWithFiles(client microClient.Client, hashes ...string) error {
    service := GetTorrentService(client)
    parameters := &proto.TorrentStrings{ Values: hashes}
    _, err := service.TorrentDeleteWithFiles(context.TODO(), parameters)
    return err
}

func TorrentReCheck(client microClient.Client, hashes ...string) error {
    service := GetTorrentService(client)
    parameters := &proto.TorrentStrings{ Values: hashes}
    _, err := service.TorrentReCheck(context.TODO(), parameters)
    return err
}

func TorrentReAnnounce(client microClient.Client, hashes ...string) error {
    service := GetTorrentService(client)
    parameters := &proto.TorrentStrings{ Values: hashes}
    _, err := service.TorrentReAnnounce(context.TODO(), parameters)
    return err
}

func TorrentIncreasePriority(client microClient.Client, hashes ...string) error {
    service := GetTorrentService(client)
    parameters := &proto.TorrentStrings{ Values: hashes}
    _, err := service.TorrentIncreasePriority(context.TODO(), parameters)
    return err
}

func TorrentDecreasePriority(client microClient.Client, hashes ...string) error {
    service := GetTorrentService(client)
    parameters := &proto.TorrentStrings{ Values: hashes}
    _, err := service.TorrentDecreasePriority(context.TODO(), parameters)
    return err
}

func TorrentSetHighestPriority(client microClient.Client, hashes ...string) error {
    service := GetTorrentService(client)
    parameters := &proto.TorrentStrings{ Values: hashes}
    _, err := service.TorrentSetHighestPriority(context.TODO(), parameters)
    return err
}

func TorrentSetLowestPriority(client microClient.Client, hashes ...string) error {
    service := GetTorrentService(client)
    parameters := &proto.TorrentStrings{ Values: hashes}
    _, err := service.TorrentSetLowestPriority(context.TODO(), parameters)
    return err
}

func TorrentSetFilePriority(client microClient.Client, hash string, priority uint64, files ...string) error {
    service := GetTorrentService(client)
    parameters := &proto.TorrentFilesPriority{ Hash: hash, Priority: priority, Files: files }
    _, err := service.TorrentSetFilePriority(context.TODO(), parameters)
    return err
}