all: gobuild gotest

godev:
	cd gogo && go run main.go

gobuild: goclean goinstall

gorebuild: goclean goreinstall

goclean:
	go clean ./...

goinstall:
	go get -t -v github.com/poseidon

goreinstall:
	go get -t -a -v github.com/poseidon

gotest:
	go test github.com/poseidon/app/controllers
	go test github.com/poseidon/app/middlewares
	# go test github.com/poseidon/app/models

gopackage:
	mkdir -p bin && go build -a -o bin/poseidon src/github.com/poseidon/main.go

travis: gobuild gotest
