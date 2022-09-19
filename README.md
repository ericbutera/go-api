# API GO

Experiment for kubernetes and golang.

## Docs

Generated with [swag](https://github.com/swaggo/swag). Install with `go install github.com/swaggo/swag/cmd/swag@latest`.

## Config

Configuration handled by `ENV` variables prefixed with `GOAPI_`. See [.env.sample](./.env.sample).

## Container Registry

- microk8s enable registry
- update Makefile version, use `make image-build` and `make image-push`.
- update kubernetes/deployment.yml
  - update image tag
  - `kubectl apply -f kubernetes/deployment.yml`

## Telemetry

Goals:

- DataDog
- OpenTelemetry

Forward:

- Jaeger: `kubectl port-forward $(kubectl get pods -l=app="jaeger" -o name) 16686:16686`
- Collector: `kubectl port-forward -n default service/simplest-collector 14268:14268 --address 0.0.0.0`

[Jaeger UI](http://localhost:16686/)

## Debug

See [DEBUG](./DEBUG.md).

## CI/CD

- github actions
- <https://github.com/nektos/act>
- <https://github.com/keel-hq/keel>
- <https://github.com/argoproj>

## Notes

- GOPATH is a project level [env var](https://github.com/travis-ci/gimme/issues/240).
- use `gimme` to manage go
- gimme has a different base path between linux and macos (with brew).
