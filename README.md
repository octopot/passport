> # 👮 Passport
>
> Person Identifier as a Service &mdash; your personal user tracker and auth service.

[![Build][build.icon]][build.page]
[![Template][template.icon]][template.page]

## Roadmap

- [x] v1: [MVP][project_v1]
  - [**January 5, 2019**][project_v1_dl]
  - Main concepts and working prototype.
- [ ] v2: [Tracking][project_v2]
  - [**Sometime, 20xx**][project_v2_dl]
  - Tracking center.
- [ ] v3: [Authentication][project_v3]
  - [**Somehow, 20xx**][project_v3_dl]
  - Authentication center.

## Motivation

- We have to full control over our users' data and protect it from third parties.

## Quick start

Requirements:

- Docker 18.06.0-ce or above
- Docker Compose 1.22.0 or above
- Go 1.9.2 or above
- GNU Make 3.81 or above

```bash
$ make up demo status

       Name                     Command               State                           Ports
-------------------------------------------------------------------------------------------------------------------
passport_db_1        docker-entrypoint.sh postgres    Up      0.0.0.0:5432->5432/tcp
passport_server_1    /bin/sh -c echo $BASIC_USE ...   Up      0.0.0.0:443->443/tcp, 0.0.0.0:80->80/tcp
passport_service_1   service run --with-profili ...   Up      0.0.0.0:8080->8080/tcp, 0.0.0.0:8090->8090/tcp,
                                                              0.0.0.0:8091->8091/tcp, 8092/tcp, 8093/tcp

$ make help
```

## Specification

### API

You can find API specification [here](env/client/rest.http). Also, we recommend using [Insomnia](https://insomnia.rest/)
HTTP client to work with the API - you can import data for it from the [file](env/client/insomnia.json).

### CLI

<details>
<summary><strong>Service command-line interface</strong></summary>

```bash
$ make install

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
# or use mirror
$ docker pull quay.io/kamilsk/passport:1.x
```

### From source code

```bash
$ egg github.com/kamilsk/passport@^1.0.0 -- make test install
$ # or use mirror
$ egg bitbucket.org/kamilsk/passport@^1.0.0 -- make test install
```

> [egg](https://github.com/kamilsk/egg)<sup id="anchor-egg">[1](#egg)</sup> is an `extended go get`.

<sup id="egg">1</sup> The project is still in prototyping.[↩](#anchor-egg)

---

made with ❤️ for everyone

[build.page]:       https://travis-ci.com/octopot/passport
[build.icon]:       https://travis-ci.com/octopot/passport.svg?branch=master
[design.page]:      https://www.notion.so/octolab/Passport-42fe8035237a445582af92546f7a48c4?r=0b753cbf767346f5a6fd51194829a2f3
[promo.page]:       https://octopot.github.io/passport/
[template.page]:    https://github.com/octomation/go-service
[template.icon]:    https://img.shields.io/badge/template-go--service-blue

[egg]:              https://github.com/kamilsk/egg

[project_v1]:       https://github.com/octopot/passport/projects/2
[project_v1_dl]:    https://github.com/octopot/passport/milestone/1
[project_v2]:       https://github.com/octopot/passport/projects/3
[project_v2_dl]:    https://github.com/octopot/passport/milestone/2
[project_v3]:       https://github.com/octopot/passport/projects/4
[project_v3_dl]:    https://github.com/octopot/passport/milestone/3
