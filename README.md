> # Passport [![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Person%20Identifier%20as%20a%20Service&url=https://kamilsk.github.io/passport/&via=ikamilsk&hashtags=go,service,authentication,identification)
> [![Analytics](https://ga-beacon.appspot.com/UA-109817251-24/passport/readme?pixel)](https://kamilsk.github.io/passport/)
> Person Identifier as a Service.

[![Patreon](https://img.shields.io/badge/patreon-donate-orange.svg)](https://www.patreon.com/octolab)
[![Build Status](https://travis-ci.org/kamilsk/passport.svg?branch=master)](https://travis-ci.org/kamilsk/passport)
[![Coverage Status](https://coveralls.io/repos/github/kamilsk/passport/badge.svg)](https://coveralls.io/github/kamilsk/passport)
[![GoDoc](https://godoc.org/github.com/kamilsk/passport?status.svg)](https://godoc.org/github.com/kamilsk/passport)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Quick start

```bash
$ make up status

    Name                   Command               State                                  Ports
-----------------------------------------------------------------------------------------------------------------------------
env_db_1        docker-entrypoint.sh postgres    Up      0.0.0.0:5432->5432/tcp
env_server_1    /bin/sh -c envsubst '$SERV ...   Up      80/tcp, 0.0.0.0:80->8080/tcp
env_service_1   passport run --with-profil ...   Up      0.0.0.0:8080->80/tcp, 0.0.0.0:8090->8090/tcp, 0.0.0.0:8091->8091/tcp
```

## Specification

### API

You can find API specification [here](env/rest.http). Also, we recommend using [Insomnia](https://insomnia.rest)
HTTP client to work with the API - you can import data for it from the [file](env/insomnia.json).

### CLI

```bash
$ passport --help
Passport

Usage:
   [command]

Available Commands:
  help        Help about any command
  migrate     Apply database migration
  run         Start HTTP server
  version     Show application version

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.
```

## Installation

### Brew

```bash
$ brew install kamilsk/tap/passport
```

### Binary

```bash
$ export VER=1.0.0      # all available versions are on https://github.com/kamilsk/passport/releases
$ export REQ_OS=Linux   # macOS and Windows are also available
$ export REQ_ARCH=64bit # 32bit is also available
$ wget -q -O passport.tar.gz \
       https://github.com/kamilsk/passport/releases/download/"${VER}/passport_${VER}_${REQ_OS}-${REQ_ARCH}".tar.gz
$ tar xf passport.tar.gz -C "${GOPATH}"/bin/ && rm passport.tar.gz
```

### Docker Hub

```bash
$ docker pull kamilsk/passport:latest
```

### From source code

```bash
$ egg github.com/kamilsk/passport@^1.0.0 -- make generate test install
```

#### Mirror

```bash
$ egg bitbucket.org/kamilsk/passport@^1.0.0 -- make generate test install
```

> [egg](https://github.com/kamilsk/egg) is an `extended go get`.

#### Requirements

- Docker 17.09.0-ce or above
- Docker Compose 1.16.1 or above
- Go 1.9.2 or above
- GNU Make 3.81 or above

## Notes

- brief roadmap
  - [x] v1: MVP
  - [ ] v2: ...
  - [ ] v3: ...
  - [ ] v4: CRUD
  - [ ] v5: GUI
  - [ ] Passport, SaaS
- tested on Go 1.9 and 1.10

---

[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/kamilsk/passport)
[![@kamilsk](https://img.shields.io/badge/author-%40kamilsk-blue.svg)](https://twitter.com/ikamilsk)
[![@octolab](https://img.shields.io/badge/sponsor-%40octolab-blue.svg)](https://twitter.com/octolab_inc)

made with ❤️ by [OctoLab](https://www.octolab.org/)
