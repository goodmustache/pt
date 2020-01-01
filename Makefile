GO_FILES = $(wildcard *.go) $(wildcard */*.go)
NODES = 4

all : pt

test : pt
	go fmt ./...
	go vet . ./actions ./commands ./config ./tracker
	ginkgo -nodes $(NODES) -r -randomizeSuites -randomizeAllSpecs -race

install:
	go install .

clean :
	rm pt

pt : $(GO_FILES)
	go build .

.PHONY : all test install clean
