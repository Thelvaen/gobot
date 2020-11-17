@echo off
echo Deleting old binary
del gobot.exe
set GOOS=windows
set GOARCH=amd64
echo Compiling Templates into assets.go
go-assets-builder includes layouts -o assets.go
echo Compiling assets into bindata.go
go-bindata -fs -prefix "static/" -o bindata.go static/...
echo Compiling the binary
go build
echo Starting the bot
gobot