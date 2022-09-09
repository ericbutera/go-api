# API GO

Experiment for kubernetes and golang.

## Docs

Generated with [swag](https://github.com/swaggo/swag). Install with `go install github.com/swaggo/swag/cmd/swag@latest`.

## Config

Configuration handled by `ENV` variables prefixed with `GOAPI_`. See [.env.sample](./.env.sample).

## Debug

See [DEBUG](./DEBUG.md).

## Notes

- GOPATH is a project level [env var](https://github.com/travis-ci/gimme/issues/240).
- use `gimme` to manage go
- gimme has a different base path between linux and macos (with brew).