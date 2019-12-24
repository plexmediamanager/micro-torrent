package service

import (
    "encoding/json"
    microClient "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-torrent/errors"
    "github.com/plexmediamanager/micro-torrent/proto"
)

// Convert response to structure
func protoToStructure(output interface{}, result *proto.TorrentResponse, err error) error {
    if err != nil {
        return err
    }
    err = json.Unmarshal(result.Result, &output)
    if err != nil {
        return errors.UnmarshalError.ToError(err)
    }
    return nil
}

func GetTorrentService(client microClient.Client) proto.TorrentService {
    return proto.NewTorrentService("micro.torrent", client)
}
