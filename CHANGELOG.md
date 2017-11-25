## 0.4.1 (2017-11-25)

change html and css

## 0.4.0 (2017-11-25)

add metadata(registeredBy, createdAt) for entries and drop os dependent commands(peco, open)

## BC break

- delete open, peco command support
- dump cli,api return new json schema which include RegisteredBy, CreatedAt

## New Feature

- add `registeredBy` to link attribute
  - for delete, same registeredBy should be passed.
  - it is *not* secure, but it will be enough for team use.
- api, cli's supports (optional) registeredBy parameter
  - add $key $link [$registeredBy]
  - curl -X POST -d "link=http://foo[&registeredBy=bar]"
  - rm $key [$registeredBy]
  - curl -X DELETE "http://localhost:5656/foo[?registeredBy=bar]"

## Internal Change

- data schema was changed.(gol will update for old schema automatically)

## 0.3.0 (2017-11-7)

add get/post/delete ui

### Breaking Change

- `GET /` returns html contents which include get/post/delete UI.
- `GET /api/dump` returns all links as json (previously, this contents was returned by / )

## 0.2.1 (2017-11-4)

add force option

### Added

- `curl -d "value=${some_url}&force=true" localhost:5656/${key}`
- `gol add --force ${key} ${some_url}`

## 0.2.0 (2017-08-26)

add post/delete api

### Added

- add post api for adding a new link
- add delete api for removing a link

### Breaking Change

- now, key can't contain "/"

## 0.1.1 (2017-06-24)

add dump

### Added

- Add dump

### Deprecated

- Nothing

### Removed

- Nothing

### Fixed

- Nothing

## 0.1.0 (2017-06-24)

Initial release

### Added

- Add get/add/rm/ls/open/peco/server

### Deprecated

- Nothing

### Removed

- Nothing

### Fixed

- Nothing
