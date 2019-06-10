#!/bin/bash

#mac 版的binary
go build  -o goPatchPrometheus *.go
mv goPatchPrometheus ../binary/mac

#Linux版的binary
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o goPatchPrometheus *.go
mv goPatchPrometheus ../binary/Linux

#Windows版的binary
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o goPatchPrometheus *.go
mv goPatchPrometheus ../binary/Windows