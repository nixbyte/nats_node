CGO_ENABLED=0
GOPATH=${CURDIR}/../..:${CURDIR}

build: linux  mac
	go build -o bin/main ./main.go

linux:
	GOOS=linux GOARCH=amd64 go build -o bin/main_linux ./main.go

mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/main_darwin ./main.go

init:
	dep init

ensure:
	dep ensure

update:
	dep ensure -update

status:
	dep status

check:
	dep check


f:
	gofmt -l -s -w `find . -type f -name '*.go' -not -path "./vendor/*"`
