# Mattermost Wake-on-LAN Plugin [![CircleCI branch](https://img.shields.io/circleci/project/github/maruTA-bis5/mattermost-plugin-wol/master.svg)](https://circleci.com/gh/maruTA-bis5/mattermost-plugin-wol)

This plugin provides the slash command to send the Wake on LAN magic packet.

## Installation
WIP

## Configuration
WIP

## Development
Note that this project uses [Go modules](https://github.com/golang/go/wiki/Modules). Be sure to locate the project outside of `$GOPATH`, or allow the use of Go modules within your `$GOPATH` with an `export GO111MODULE=on`.

### Build
```
make
```

This will produce a single plugin file (with support for multiple architectures) for upload to your Mattermost server:

```
dist/net.bis5.mattermost.plugins.wol.tar.gz
```

There is a build target to automate deploying and enabling the plugin to your server, but it requires configuration and [http](https://httpie.org/) to be installed:
```
export MM_SERVICESETTINGS_SITEURL=http://localhost:8065
export MM_ADMIN_USERNAME=admin
export MM_ADMIN_PASSWORD=password
make deploy
```

Alternatively, if you are running your `mattermost-server` out of a sibling directory by the same name, use the `deploy` target alone to  unpack the files into the right directory. You will need to restart your server and manually enable your plugin.

In production, deploy and upload your plugin via the [System Console](https://about.mattermost.com/default-plugin-uploads).

## License
Apache License, Version 2.0

## Development
### Generate Changelog
WIP
```
go get github.com/Songmu/ghch/cmd/ghch
ghch --format=markdown -w --next-version=vX.Y.Z
```
