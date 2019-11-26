VERSION := $(shell git rev-parse --short HEAD)
install:
	go install -ldflags '-X main.VersionString=$(VERSION)'
