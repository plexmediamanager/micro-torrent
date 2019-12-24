package resolver

import (
    "encoding/json"
    "github.com/plexmediamanager/micro-torrent/errors"
    "github.com/plexmediamanager/micro-torrent/proto"
    "github.com/plexmediamanager/micro-torrent/torrent"
)

type TorrentService struct {
    Torrent           *torrent.Client
}

// Convert structure to bytes and return error if needed
func structureToBytesWithError(structure interface{}, err error, response *proto.TorrentResponse) error {
    if err != nil {
        return err
    }
    bytes, err := json.Marshal(structure)
    if err != nil {
        return errors.UnmarshalError.ToError(err)
    }
    response.Result = bytes
    return nil
}

