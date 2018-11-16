BINARY = armour
GOARCH = amd64

RELEASE?=v0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date '+%Y-%m-%d_%H:%M:%S')

CURRENT_DIR = $(shell pwd)
RELEASE_DIR = ${CURRENT_DIR}/data/release

LDFLAGS = -ldflags "-X main.Release=${RELEASE} -X main.Commit=${COMMIT} -X main.BuildTime=${BUILD_TIME}"

.PHONY: default
default:
	go build ${LDFLAGS} -o ${BINARY} .

.PHONY: release
release:
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${RELEASE_DIR}/${BINARY}-linux-${GOARCH} .

.PHONY: clean
clean:
	-rm -f ${RELEASE_DIR}/${BINARY}-linux-${GOARCH}

