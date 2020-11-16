@echo off
go-assets-builder includes layouts -o assets.go
go build
gobot