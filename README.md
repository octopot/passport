> # 👮 Passport [![Tweet][icon_twitter]][twitter_publish] <img align="right" width="126" src=".github/character.png">
> [![Analytics][analytics_pixel]][page_promo]
> Person Identifier as a Service &mdash; your personal user tracker and auth service.

[![Patreon][icon_patreon]](https://www.patreon.com/octolab)
[![Build Status][icon_build]][page_build]
[![Code Coverage][icon_coverage]][page_quality]
[![Code Quality][icon_quality]][page_quality]
[![Research][icon_research]](../../tree/research)
[![License][icon_license]](LICENSE)

## Roadmap

- [ ] v1: [MVP][project_v1]
  - [**September 30, 2018**][project_v1_dl]
  - Main concepts and working prototype.
- [ ] v2: [Notification][project_v2]
  - [**November 30, 2018**][project_v2_dl]
  - Notification center.
- [ ] v3: [Authentication][project_v3]
  - [**December 31, 2018**][project_v3_dl]
  - Authentication center.

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

You can find API specification [here](env/client/rest.http). Also, we recommend using [Insomnia](https://insomnia.rest/)
HTTP client to work with the API - you can import data for it from the [file](env/client/insomnia.json).

### CLI

<details>
<summary><strong>Service command-line interface</strong></summary>

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
</details>

#### Bash and Zsh completions

You can find completion files [here](https://github.com/kamilsk/shared/tree/dotfiles/bash_completion.d) or
build your own using these commands

```bash
$ passport completion bash > /path/to/bash_completion.d/passport.sh
$ passport completion zsh  > /path/to/zsh-completions/_passport.zsh
```

## Installation

### Brew

```bash
$ brew install kamilsk/tap/passport
```

### Binary

```bash
$ export REQ_VER=1.0.0  # all available versions are on https://github.com/kamilsk/passport/releases
$ export REQ_OS=Linux   # macOS and Windows are also available
$ export REQ_ARCH=64bit # 32bit is also available
$ # wget -q -O passport.tar.gz
$ curl -sL -o passport.tar.gz \
       https://github.com/kamilsk/passport/releases/download/"${REQ_VER}/passport_${REQ_VER}_${REQ_OS}-${REQ_ARCH}".tar.gz
$ tar xf passport.tar.gz -C "${GOPATH}"/bin/ && rm passport.tar.gz
```

### Docker Hub

```bash
$ docker pull kamilsk/passport:1.x
```

### From source code

```bash
$ egg github.com/kamilsk/passport@^1.0.0 -- make test install
$ # or use mirror
$ egg bitbucket.org/kamilsk/passport@^1.0.0 -- make test install
```

> [egg](https://github.com/kamilsk/egg)<sup id="anchor-egg">[1](#egg)</sup> is an `extended go get`.

## Update

This application is in a state of [MVP](https://en.wikipedia.org/wiki/Minimum_viable_product) and under active
development. [SemVer](https://semver.org/) is used for releases, and you can easily be updated within minor versions,
but major versions can be not [BC](https://en.wikipedia.org/wiki/Backward_compatibility)-safe.

<sup id="egg">1</sup> The project is still in prototyping. [↩](#anchor-egg)

---

[![Gitter][icon_gitter]](https://gitter.im/kamilsk/passport)
[![@kamilsk][icon_tw_author]](https://twitter.com/ikamilsk)
[![@octolab][icon_tw_sponsor]](https://twitter.com/octolab_inc)

made with ❤️ by [OctoLab](https://www.octolab.org/)

[analytics_pixel]: https://ga-beacon.appspot.com/UA-109817251-24/passport/readme?pixel

[icon_build]:      https://travis-ci.org/kamilsk/passport.svg?branch=master
[icon_coverage]:   https://scrutinizer-ci.com/g/kamilsk/passport/badges/coverage.png?b=master
[icon_gitter]:     https://badges.gitter.im/Join%20Chat.svg
[icon_license]:    https://img.shields.io/badge/license-MIT-blue.svg
[icon_patreon]:    https://img.shields.io/badge/patreon-donate-orange.svg
[icon_quality]:    https://scrutinizer-ci.com/g/kamilsk/passport/badges/quality-score.png?b=master
[icon_research]:   https://img.shields.io/badge/research-in%20progress-yellow.svg
[icon_tw_author]:  https://img.shields.io/badge/author-%40kamilsk-blue.svg
[icon_tw_sponsor]: https://img.shields.io/badge/sponsor-%40octolab-blue.svg
[icon_twitter]:    https://img.shields.io/twitter/url/http/shields.io.svg?style=social

[page_build]:      https://travis-ci.org/kamilsk/passport
[page_promo]:      https://kamilsk.github.io/passport/
[page_quality]:    https://scrutinizer-ci.com/g/kamilsk/passport/?branch=master

[project_v1]:      https://github.com/kamilsk/passport/projects/2
[project_v1_dl]:   https://github.com/kamilsk/passport/milestone/1
[project_v2]:      https://github.com/kamilsk/passport/projects/3
[project_v2_dl]:   https://github.com/kamilsk/passport/milestone/2
[project_v3]:      https://github.com/kamilsk/passport/projects/4
[project_v3_dl]:   https://github.com/kamilsk/passport/milestone/3

[twitter_publish]: https://twitter.com/intent/tweet?text=Person%20Identifier%20as%20a%20Service&url=https://kamilsk.github.io/passport/&via=ikamilsk&hashtags=go,service,authentication,identification
