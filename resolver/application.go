package resolver

import (
    "context"
    "github.com/plexmediamanager/micro-torrent/proto"
)

func (service TorrentService) ApplicationVersion (_ context.Context, properties *proto.TorrentEmpty, response *proto.TorrentResponse) error {
    result, err := service.Torrent.ApplicationVersion()
    return structureToBytesWithError(result, err, response)
}

func (service TorrentService) ApplicationAPIVersion (_ context.Context, properties *proto.TorrentEmpty, response *proto.TorrentResponse) error {
    result, err := service.Torrent.ApplicationAPIVersion()
    return structureToBytesWithError(result, err, response)
}

func (service TorrentService) ApplicationBuildInformation (_ context.Context, properties *proto.TorrentEmpty, response *proto.TorrentResponse) error {
    result, err := service.Torrent.ApplicationBuildInformation()
    return structureToBytesWithError(result, err, response)
}

func (service TorrentService) ApplicationPreferences (_ context.Context, properties *proto.TorrentEmpty, response *proto.TorrentResponse) error {
    result, err := service.Torrent.ApplicationPreferences()
    return structureToBytesWithError(result, err, response)
}

