# Gobot
This is a full featured application (still in progress), hence, it is not intended to be imported in your project.

# Twitch Authentification
To get your oauth token, please go to https://twitchapps.com/tmi/ and click connect, then Copy/Paste your token in the Token field of the config.yml

# Configuration tips
* If Port is not defined in the config, it will use 8090 as default,
* Default stack size for aggregator is 60,
* Cred/Channel must always be defined, as it gives the bot its main Channel for operation (as well as the Twitch ID to use when using oAuth),

# Functions
* user authentication, initial user must be provided in init.yml (broken since reworked to IRIS, not yet entirely fixed, and mliddleware to restrict access is not yet done),
* Aggregated messages (AggregChans array in the config.yml), go to http://[server]:Port/auth/messages
* Dice rolling (!dice/!dice XX/!rand XX)
* Score for sentences that are longer than 10 chars (!score) or /auth/stats

# Before you compile
you need GCC installed on your system, as well as sqlite.so or sqlite.dll in your library path
then you will need to run the following commands (if you never ran them on your system
```
go get github.com/mattn/go-sqlite3
go install github.com/mattn/go-sqlite3
```

# Additional tools required :
go-bindata need to be installed with the following commands (the ... are important)
```
go get -u github.com/go-bindata/go-bindata/...
```

# Compilation
First build templates with (the ... are important):
```
go-bindata -fs -pkg templates -o templates/templates.go -prefix "html/templates/" html/templates/...
```

Then build static assets with (the ... are important):
```
go-bindata -fs -pkg static -o static/static.go -prefix "html/static/" html/static/...
```

Then you can build the bot with:
```
go build
```

# Runing in background (systemd)
copy the gobot binary to /usr/sbin, then create a service file in /etc/systemd/system/ :
```
sudo nano /etc/systemd/system/gobot.service
```
```
[Unit]
Description=Twitch Bot
After=network.target

[Service]
ExecStart=/usr/sbin/gobot

[Install]
WantedBy=multi-user.target
```

then execute:
```
sudo systemctl daemon-reload
sudo systemctl enable gobot
sudo systemctl start gobot
```
