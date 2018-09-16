> # Passport [![Tweet][icon_twitter]][publish_twitter] <img align="right" width="100" height="100" src=".github/character.png">
> [![Analytics](https://ga-beacon.appspot.com/UA-109817251-24/passport/readme?pixel)](https://kamilsk.github.io/passport/)
> Person Identifier as a Service &mdash; your personal user tracker and auth service.

[![Patreon](https://img.shields.io/badge/patreon-donate-orange.svg)](https://www.patreon.com/octolab)
[![Build Status](https://travis-ci.org/kamilsk/passport.svg?branch=master)](https://travis-ci.org/kamilsk/passport)
[![Code Coverage](https://scrutinizer-ci.com/g/kamilsk/passport/badges/coverage.png?b=master)](https://scrutinizer-ci.com/g/kamilsk/passport/?branch=master)
[![Code Quality](https://scrutinizer-ci.com/g/kamilsk/passport/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/kamilsk/passport/?branch=master)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Quick start

Requirements: 

- Docker 17.09.0-ce or above
- Docker Compose 1.16.1 or above
- Go 1.9.2 or above
- GNU Make 3.81 or above

```bash
$ make up status

       Name                     Command               State                                  Ports
----------------------------------------------------------------------------------------------------------------------------------
passport_db_1        docker-entrypoint.sh postgres    Up      0.0.0.0:5432->5432/tcp
passport_server_1    /bin/sh -c envsubst '$SERV ...   Up      80/tcp, 0.0.0.0:80->8080/tcp
passport_service_1   passport run --with-profil ...   Up      0.0.0.0:8080->80/tcp, 0.0.0.0:8090->8090/tcp, 0.0.0.0:8091->8091/tcp

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
  passport [command]

Available Commands:
  completion  Print Bash or Zsh completion
  help        Help about any command
  migrate     Apply database migration
  run         Start HTTP server
  version     Show application version

Flags:
  -h, --help   help for passport

Use "passport [command] --help" for more information about a command.
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
$ docker pull kamilsk/passport:1.x
```

### From source code

```bash
$ egg github.com/kamilsk/passport@^1.0.0 -- make test install
```

#### Mirror

```bash
$ egg bitbucket.org/kamilsk/passport@^1.0.0 -- make test install
```

> [egg](https://github.com/kamilsk/egg) is an `extended go get`.

### Bash and Zsh completions

You can find completion files [here](https://github.com/kamilsk/shared/tree/dotfiles/bash_completion.d) or
build your own using these commands

```bash
$ passport completion bash > /path/to/bash_completion.d/passport.sh
$ passport completion zsh  > /path/to/zsh-completions/_passport.zsh
```

## Notes

- brief roadmap
  - [x] v1: MVP
  - [ ] v2: Simple Auth
  - [ ] v3: OAuth
  - [ ] v4: CRUD
  - [ ] v5: GUI
  - [ ] Passport, SaaS
- [research](../../tree/research)
- tested on Go 1.9 and 1.10

### Update

This application is in a state of [MVP](https://en.wikipedia.org/wiki/Minimum_viable_product) and under active
development. [SemVer](https://semver.org/) is used for releases, and you can easily be updated within minor versions,
but major versions can be not [BC](https://en.wikipedia.org/wiki/Backward_compatibility)-safe.

---

[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/kamilsk/passport)
[![@kamilsk](https://img.shields.io/badge/author-%40kamilsk-blue.svg)](https://twitter.com/ikamilsk)
[![@octolab](https://img.shields.io/badge/sponsor-%40octolab-blue.svg)](https://twitter.com/octolab_inc)

made with ❤️ by [OctoLab](https://www.octolab.org/)

[icon_twitter]:    https://img.shields.io/twitter/url/http/shields.io.svg?style=social

[publish_twitter]: https://twitter.com/intent/tweet?text=Person%20Identifier%20as%20a%20Service&url=https://kamilsk.github.io/passport/&via=ikamilsk&hashtags=go,service,authentication,identification
