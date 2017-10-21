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


deploy:
	GOOS=linux GOARCH=amd64 go build -o bin/poseidon gogo/main.go
	scp bin/poseidon root@115.231.180.69:/root/poseidon/bin/poseidon2
	ssh root@115.231.180.69 "supervisorctl stop poseidon && mv /root/poseidon/bin/poseidon2 /root/poseidon/bin/poseidon && supervisorctl start poseidon"
