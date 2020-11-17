# gobot
To get your oauth token, please go to https://twitchapps.com/tmi/ and click connect, then Copy/Paste your token in the Token field of the config.yml

If Port is not defined in the config, it will use 8090 as default.

# Functions
* Aggregated messages (AggregChans array in the config.yml), go to http://[server]:Port/messages

# Before you compile
you need GCC installed on your system, as well as sqlite.so or sqlite.dll in your library path
then you will need to run the following commands (if you never ran them on your system
```
go get github.com/mattn/go-sqlite3
go install github.com/mattn/go-sqlite3
```

# Additional tools required :
go-bindata & go-assets-builder need to be installed with the following commands (the ... are important)
```
go get -u github.com/go-bindata/go-bindata/...
go get github.com/jessevdk/go-assets-builder
```

# Compilation
First build templates with:
```
go-assets-builder includes layouts -o assets.go
```

Then build static assets with (the ... are important):
```
go-bindata -fs -prefix "static/" -o bindata.go static/...
```

Then you can build the bot with:
```
go build
```