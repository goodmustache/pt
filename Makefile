GO_FILES = $(wildcard *.go) $(wildcard */*.go)

all : pt

test : pt
	go fmt ./...
	go vet . ./commands ./config
	ginkgo -r

clean :
	rm pt

pt : $(GO_FILES)
	go build -o pt main.go

.PHONY : all test clean
