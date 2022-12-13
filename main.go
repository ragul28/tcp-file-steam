package main

import (
	"time"

	"github.com/ragul28/tcp-file-stream/pkg/client"
	"github.com/ragul28/tcp-file-stream/pkg/fileserver"
)

func main() {

	go func() {
		time.Sleep(2 * time.Second)
		client.SendFile(4000)
	}()

	server := &fileserver.FileServer{}
	server.Start()
}
