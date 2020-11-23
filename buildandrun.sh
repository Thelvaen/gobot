#!/usr/bin/env bash
CPU=`lscpu | head -n 1 | awk '{print $2}'`
if [ -f gobot ]; then
    echo Deleting old binary
    rm gobot
fi
echo Configuring env for compilation
set GOOS=linux
if [ $CPU = "aarch64" ]; then
    set GOARCH=arm64
else
    set GOARCH=amd64
fi
echo Building templates file
go-bindata -fs -pkg templates -o templates/templates.go -prefix "html/templates/" html/templates/...
echo Building static assets file
go-bindata -fs -pkg static -o static/static.go -prefix "html/static/" html/static/...
echo Compiling the bot
time go build
echo Starting the bot as root to get access to certs
sudo ./gobot