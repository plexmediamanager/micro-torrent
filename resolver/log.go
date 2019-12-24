package resolver

import (
    "context"
    "github.com/plexmediamanager/micro-torrent/proto"
)

func (service TorrentService) LogEntries (_ context.Context, properties *proto.TorrentOptions, response *proto.TorrentResponse) error {
    result, err := service.Torrent.LogEntries(properties.Options)
    return structureToBytesWithError(result, err, response)
}

