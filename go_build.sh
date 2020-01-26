#!/bin/bash

echo "build mac os version"
go build -o builds/valera_macos -v *.go
echo "build success"

echo ""
echo "build windows version"
GOOS=windows GOARCH=386 go build -o builds/valera.exe -v *.go
echo "build success"

echo ""
echo "build linux version"
GOOS=linux GOARCH=amd64 go build -o builds/valera_lunux -v *.go
echo "build success"

echo ""
echo "generate test coverage"
gopherbadger -md="README.md"