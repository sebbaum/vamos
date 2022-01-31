BINARY_NAME=Vamos
VERSION=$(shell git describe --always --tags)
DESTINATION=./../${BINARY_NAME}
FLAGS=-X github.com/sebbaum/vamos/cmd.version=${VERSION}

.PHONY: build
build: clean
	go build -ldflags "${FLAGS}" -o ${DESTINATION} .

.PHONY: clean
clean:
	rm -f ./../Vamos