.PHONY: compile

all: compile

compile:
	cp ./bin/sres /usr/local/bin
	go build -o ./bin/gap ./app/gap/main.go