package main

import (
	"net"
)

type FileServer struct{}

func (fs *FileServer) start() {

}

func (fs *FileServer) readLoop(conn net.Conn) {

}

func main() {

	server := &FileServer{}
	server.start()
}
