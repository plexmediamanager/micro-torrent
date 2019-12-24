package main

import (
    format "fmt"
    "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-torrent/proto"
    "github.com/plexmediamanager/micro-torrent/resolver"
    "github.com/plexmediamanager/micro-torrent/torrent"
    "github.com/plexmediamanager/service"
    "github.com/plexmediamanager/service/log"
    "time"
)

func main() {
    application := service.CreateApplication()

    err := application.InitializeConfiguration()
    if err != nil {
        log.Panic(err)
    }

    err = application.InitializeMicroService()
    if err != nil {
        log.Panic(err)
    }

    err = application.Service().Client().Init(
        client.PoolSize(10),
        client.Retries(30),
        client.RequestTimeout(1 * time.Second),
    )
    if err != nil {
        log.Panic(err)
    }

    torrentClient := torrent.Initialize()
    err = torrentClient.Authenticate()
    if err != nil {
        format.Println(err)
    }

    err = proto.RegisterTorrentServiceHandler(application.Service().Server(), resolver.TorrentService{ Torrent: torrentClient })
    if err != nil {
        log.Panic(err)
    }

    go application.StartMicroService()

    service.WaitForOSSignal(1)
}
