CGO_ENABLED=0
GOPATH=${CURDIR}/../..:${CURDIR}

filename=swagger
	

build: build_server build_client

build_server: linux_server  mac_server
	go build -o bin/current_server ./server.go

linux_server:
	GOOS=linux GOARCH=amd64 go build -o bin/server_linux ./server.go

mac_server:
	GOOS=darwin GOARCH=amd64 go build -o bin/server_darwin ./server.go

build_client: linux_client  mac_client
	go build -o bin/current_client ./client.go

linux_client:
	GOOS=linux GOARCH=amd64 go build -o bin/client_linux ./client.go

mac_client:
	GOOS=darwin GOARCH=amd64 go build -o bin/client_darwin ./client.go

init:
	go mod init

ensure:
	go mod vendor

update:
	go mod tidy

check:
	go mod verify

check_swagger:
	which swagger || (go get -u  github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check_swagger
	swagger generate spec -o ./docs/swagger/$(FILENAME).json -c $(FILENAME) --scan-models

run_swagger: check_swagger
	swagger serve -F=swagger ./docs/swagger/$(FILENAME).json


	gofmt -l -s -w `find . -type f -name '*.go' -not -path "./vendor/*"`
