package resolver

import (
    "context"
    "github.com/plexmediamanager/micro-torrent/proto"
)

func (service TorrentService) TorrentList (_ context.Context, parameters *proto.TorrentEmpty, response *proto.TorrentResponse) error {
    result, err := service.Torrent.TorrentList()
    return structureToBytesWithError(result, err, response)
}

func (service TorrentService) TorrentProperties (_ context.Context, parameters *proto.TorrentString, response *proto.TorrentResponse) error {
    result, err := service.Torrent.TorrentProperties(parameters.Value)
    return structureToBytesWithError(result, err, response)
}

func (service TorrentService) TorrentTrackers (_ context.Context, parameters *proto.TorrentString, response *proto.TorrentResponse) error {
    result, err := service.Torrent.TorrentTrackers(parameters.Value)
    return structureToBytesWithError(result, err, response)
}

func (service TorrentService) TorrentContents (_ context.Context, parameters *proto.TorrentString, response *proto.TorrentResponse) error {
    result, err := service.Torrent.TorrentContents(parameters.Value)
    return structureToBytesWithError(result, err, response)
}

func (service TorrentService) TorrentPiecesState (_ context.Context, parameters *proto.TorrentString, response *proto.TorrentResponse) error {
    result, err := service.Torrent.TorrentPiecesState(parameters.Value)
    return structureToBytesWithError(result, err, response)
}

func (service TorrentService) TorrentPiecesHashes (_ context.Context, parameters *proto.TorrentString, response *proto.TorrentResponse) error {
    result, err := service.Torrent.TorrentPiecesHashes(parameters.Value)
    return structureToBytesWithError(result, err, response)
}

func (service TorrentService) TorrentPause (_ context.Context, parameters *proto.TorrentStrings, response *proto.TorrentEmpty) error {
    return service.Torrent.TorrentPause(parameters.Values...)
}

func (service TorrentService) TorrentResume (_ context.Context, parameters *proto.TorrentStrings, response *proto.TorrentEmpty) error {
    return service.Torrent.TorrentResume(parameters.Values...)
}

func (service TorrentService) TorrentDelete (_ context.Context, parameters *proto.TorrentStrings, response *proto.TorrentEmpty) error {
    return service.Torrent.TorrentDelete(parameters.Values...)
}

func (service TorrentService) TorrentDeleteWithFiles (_ context.Context, parameters *proto.TorrentStrings, response *proto.TorrentEmpty) error {
    return service.Torrent.TorrentDeleteWithFiles(parameters.Values...)
}

func (service TorrentService) TorrentReCheck (_ context.Context, parameters *proto.TorrentStrings, response *proto.TorrentEmpty) error {
    return service.Torrent.TorrentReCheck(parameters.Values...)
}

func (service TorrentService) TorrentReAnnounce (_ context.Context, parameters *proto.TorrentStrings, response *proto.TorrentEmpty) error {
    return service.Torrent.TorrentReAnnounce(parameters.Values...)
}

func (service TorrentService) TorrentIncreasePriority (_ context.Context, parameters *proto.TorrentStrings, response *proto.TorrentEmpty) error {
    return service.Torrent.TorrentIncreasePriority(parameters.Values...)
}

func (service TorrentService) TorrentDecreasePriority (_ context.Context, parameters *proto.TorrentStrings, response *proto.TorrentEmpty) error {
    return service.Torrent.TorrentDecreasePriority(parameters.Values...)
}

func (service TorrentService) TorrentSetHighestPriority (_ context.Context, parameters *proto.TorrentStrings, response *proto.TorrentEmpty) error {
    return service.Torrent.TorrentSetHighestPriority(parameters.Values...)
}

func (service TorrentService) TorrentSetLowestPriority (_ context.Context, parameters *proto.TorrentStrings, response *proto.TorrentEmpty) error {
    return service.Torrent.TorrentSetLowestPriority(parameters.Values...)
}

func (service TorrentService) TorrentSetFilePriority (_ context.Context, parameters *proto.TorrentFilesPriority, response *proto.TorrentEmpty) error {
    return service.Torrent.TorrentSetFilePriority(parameters.Hash, parameters.Priority, parameters.Files...)
}
