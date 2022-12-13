package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/ragul28/tcp-file-stream/pkg/client"
	"github.com/ragul28/tcp-file-stream/pkg/fileserver"
)

func main() {

	listenPort := flag.Int("l", 3000, "Server listen port")
	flag.Parse()

	go func() {
		time.Sleep(2 * time.Second)
		client.SendFile(4000)
	}()

	server := &fileserver.FileServer{}
	fmt.Println(*listenPort)
	server.Start(strconv.Itoa(*listenPort))
}
