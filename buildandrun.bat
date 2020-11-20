@echo off
echo Deleting old binary
del gobot.exe
set GOOS=windows
set GOARCH=amd64
echo Building templates file
go-bindata -fs -pkg templates -o templates/templates.go -prefix "html/templates/" html/templates/...
echo Building static assets file
go-bindata -fs -pkg static -o static/static.go -prefix "html/static/" html/static/...
echo Compiling the binary
go build
echo Starting the bot
gobot