#!/bin/sh

go get github.com/mdirkse/i3ipc
go build -o i3altlayout -ldflags="-s -w" main.go

sudo mv i3altlayout /usr/bin/i3altlayout
