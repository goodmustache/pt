GO_FILES = $(wildcard *.go) $(wildcard */*.go)
NODES ?= 4

all : pt

fakes :
	go generate ./...

test : pt fakes
	go fmt ./...
	go vet . ./actions ./commands ./config ./tracker
	ginkgo -nodes $(NODES) -r -randomizeSuites -randomizeAllSpecs -race

install :
	go install .

clean :
	rm pt

pt : $(GO_FILES)
	go build .

.PHONY : all fakes test install clean
