# gol

gol(go link) is a url shortner that run in cli and browser.

[![Build Status](https://travis-ci.org/matsu-chara/gol.svg?branch=master)](https://travis-ci.org/matsu-chara/gol)

## Usage

### in cmd

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

$ gol open myteamdocs
# open in browser (need open command)
$ gol peco jenkins
# search jenkins* and select in peco and open in browser (need peco and open command)
```

### in browser

```bash
$ gol server

# run in background
$ nohup gol server&
```

Open Chrome > Preference > add Search Engine > add below

- name: gol
- keyword: gol
- query: http://localhost:5656/%s

then, you can open gol links by
click url bar > type 'gol' > tab > type key > enter


## Install

To install, use `go get`:

```bash
$ go get github.com/matsu-chara/gol
```
