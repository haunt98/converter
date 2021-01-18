# CHANGELOG

## v0.5.0 (2021-1-18)

### Added

- feat(converter): remove prefix, postfix _

### Others

- chore: bump golangci-lint v1.34 in github action

- chore(readme): update badges

- test(converter): unit test for prefix, postfix _

- chore(changelog): generate v0.4.0

## v0.4.0 (2020-12-23)

### Added

- feat(converter): use whitelist approach

### Fixed

- fix(converter): add normalized char to check rule

### Others

- test(converter): use maxFuzzyTest for const

- refactor(converter): rename isNormalRune to isNormalChar

- refactor(converter): finalText rename to tempText

- test(converter): use fuzzy testing

- chore: add build using gotip in github action

- chore: generate CHANGELOG v0.3.0

## v0.3.0 (2020-11-13)

### Added

- feat: remove duplicate _

### Others

- chore: generate CHANGELOG v0.2.0

## v0.2.0 (2020-11-13)

### Added

- feat: add colon to deny chars

### Others

- test: add unit test for string contains tab, newline

- refactor: use deny chars as a list to remove hardcode

- chore: add MIT LICENSE

- chore: add github action badge in README

- chore: generate CHANGELOG v0.1.0

## v0.1.0 (2020-11-13)

### Added

- feat: replace ,() with _

- feat: replace ", ' with _

- feat: convert with argument

### Others

- docs: add install, usage guide in README

- refactor: move main.go to root dir for go get

- chore: add github action for go test, lint

- test: convert unit test

- refactor: move back main.go to cmd and split converter in pkg

- refactor: move main.go outta cmd
