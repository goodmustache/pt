GO_FILES = $(wildcard *.go) $(wildcard */*.go)

all : pt

test : pt
	go fmt ./...
	go vet . ./actions ./commands ./config ./tracker
	ginkgo -r -randomizeSuites -randomizeAllSpecs -race

install:
	go install .

clean :
	rm pt

pt : $(GO_FILES)
	go build -o pt main.go

.PHONY : all test install clean
