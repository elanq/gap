.PHONY: test compile

all: test compile

compile:
	cp ./bin/sres /usr/local/bin
	go build -o ./bin/gap ./app/gap/main.go

test:
	go test ./...