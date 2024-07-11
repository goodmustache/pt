GO_FILES = $(wildcard *.go) $(wildcard */*.go)
NODES ?= 4

all : test build

build: pt

fakes :
	go generate ./...

test : pt fakes
	go fmt ./...
	go vet $$(go list ./... | grep -v tools)
	go run github.com/onsi/ginkgo/ginkgo -r -randomizeSuites -randomizeAllSpecs -race

install :
	go install .

clean :
	rm pt

pt : $(GO_FILES)
	go build .

.PHONY : all fakes test install clean
