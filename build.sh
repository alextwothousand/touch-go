#!/bin/bash

if [ -d "/builds" ] 
then
	rm -rf builds
fi

## Build for linux..
GOOS=linux GOARCH=amd64 go build -o builds/linux/touch64 main.go
GOOS=linux GOARCH=386 go build -o builds/linux/touch32 main.go 

## Build for macOS.
GOOS=darwin GOARCH=amd64 go build -o builds/darwin/touch64 main.go

## Build for Windows.
GOOS=windows GOARCH=amd64 go build -o builds/win/touch64.exe main.go
GOOS=windows GOARCH=386 go build -o builds/win/touch32.exe main.go