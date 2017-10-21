all: gobuild gotest

godev:
	cd gogo && go run main.go

gobuild: goclean goinstall

gorebuild: goclean goreinstall

goclean:
	go clean ./...

goinstall:
	go get -u -v github.com/dolab/gogo
	go get github.com/dolab/session
	go get github.com/dolab/httptesting
	go get -u -v github.com/qiniu/bytes/seekable

goreinstall:
	go get -u -v github.com/dolab/gogo
	go get -t -a -v github.com/dolab/session
	go get -t -a -v github.com/dolab/httptesting
	go get -t -a -v github.com/qiniu/bytes/seekable

gotest:
	go test github.com/poseidon/app/controllers
	go test github.com/poseidon/app/middlewares

gopackage:
	mkdir -p bin && go build -a -o bin/poseidon src/github.com/poseidon/main.go

travis: gobuild gotest
