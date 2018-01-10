# gol

gol(go link) is a url shortner that run in cli and browser.

[![Build Status](https://travis-ci.org/matsu-chara/gol.svg?branch=master)](https://travis-ci.org/matsu-chara/gol)
[![Go Report Card](https://goreportcard.com/badge/github.com/matsu-chara/gol)](https://goreportcard.com/report/github.com/matsu-chara/gol)

## Install

To install, use `go get` or [dockerhub](https://hub.docker.com/r/matsuchara/gol/)

```bash
$ go get github.com/matsu-chara/gol
```

## Usage

you can use gol as a server or cli

### server

```bash
# run server at localhost:5656
$ gol server

# or run in background and restart always
$ docker-compose up -d
```

`http://localhost:5656/` will show link list and some form for registering or deleting.

#### API

##### dump

access `http://localhost:5656/api/dump` dumps all links.

![sample/gol_chrome1.png](sample/gol_chrome1.png)

##### get

access `http://localhost:5656/${key}` will return redirect response destinating to the `${value}`

##### post

`curl -d value=${some_url} http://localhost:5656/${key}"` will add a link to key

`curl -d "value=${some_url}&force=true" http://localhost:5656/${key}` will add a link to key (if replace key when conflict)

`curl -d "value=${some_url}&registeredBy=bar" http://localhost:5656/${key}` will add a link to key with a registeredBy "bar".

##### delete

`curl -X DELETE http://localhost:5656/${key}` delete a link

`curl -X DELETE http://localhost:5656/${key}?registeredBy=bar` delete a link which have registerdBy=bar. (if registeredBy is not equal to the server data, request will fail.)

#### setting chrome

Open Chrome > Preference > add Search Engine > add below

```
- name: gol
- keyword: gol
- query: http://localhost:5656/%s
```

then, you can open gol links by
click url bar > type 'gol' > tab > type key > enter

type `gol` >  tab > type key
![sample/gol_chrome2.png](sample/gol_chrome2.png)

then enter > jump to the link
![sample/gol_chrome3.png](sample/gol_chrome3.png)


As a further advanced usage, you can set custom name space.

example 1. env specific links.

```
- name: gop
- keyword: gop
- query: http://localhost:5656/production_%s
```

example 2. my original links. (prevent contamination of the global namespace)

```
- name: gom
- keyword: gom
- query: http://localhost:5656/matsu_chara_%s
```

### in cli

![sample/gol_cli.gif](sample/gol_cli.gif)

#### usage

```bash
$ gol add confluence https://confluence.nice-company.com
$ gol add myproduct https://github.nice-company.com/myteam/myproduct
$ gol add myteamdocs https://confluence.nice-company.com/pages/viewpage.action?pageId=xxxxxxx
$ gol add jenkins_dev https://dev.jenkins.nice-company.com/
$ gol add jenkins_prod https://dev.jenkins.nice-company.com/
$ gol add consul_prod https://our-server01:8500/
$ gol add myproduct_admin_prod https://our-server02:9534

$ gol ls
confluence: https://confluence.nice-company.com
consul_prod: https://our-server01:8500/
jenkins_dev: https://dev.jenkins.nice-company.com/
jenkins_prod: https://dev.jenkins.nice-company.com/
myproduct: https://github.nice-company.com/myteam/myproduct
myproduct_admin_prod: https://our-server02:9534
myteamdocs: https://confluence.nice-company.com/pages/viewpage.action?pageId=xxxxxxx

$ gol get jenkins_dev
https://dev.jenkins.nice-company.com/

$ gol rm consul_prod
$ gol get consul_prod
```

#### zsh completion

https://gist.github.com/3tty0n/0ef541bb9fce758c4c064ce96ba83a91
