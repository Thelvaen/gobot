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
go-assets-builder includes layouts -o assets.go
echo Building static assets file
go-bindata -fs -prefix "static/" -o bindata.go static/...
echo Compiling the bot
time go build
echo Starting the bot
./gobot