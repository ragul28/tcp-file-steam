build:
	GOARCH=amd64 go build -ldflags="-s -w"

install:
	go install
	
mod_init:
	go mod init github.com/ragul28/tcp-file-stream
	go get -u

mod:
	go mod tidy
	go mod verify
	go mod vendor
