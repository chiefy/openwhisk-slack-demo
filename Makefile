GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)
BINARY_NAME := helloslack

default: bin/$(BINARY_NAME)

bin:
	@mkdir -p $@

bin/$(BINARY_NAME): bin $(GOFILES)
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $@ .

.PHONY: docker-build
docker-build:
	TAG=latest
	docker build -t chiefy/openwhisk-demo:$TAG .
	docker login -u ${QUAY_USER} -p ${QUAY_TOKEN} https://quay.io
	docker push chiefy/openwhisk-demo:$TAG


