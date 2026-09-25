# Go starter

[![Deploy on velixir](https://velixir.net/img/deploy-on-velixir.svg)](https://velixir.net/new?template=go-http)

A single `main.go` on the standard library, compiled server-side into a small container.
No module dependencies to resolve, so the build is about as fast as a Go build gets.

[Deploy it on velixir](https://velixir.net/new?template=go-http).

## Running it locally

```bash
go run main.go
```

Then open http://localhost:8080.

## Deploying

```bash
velixir deploy
```

velixir compiles your source on the build node, so a cold build never ties up your laptop
and there is no Dockerfile to maintain.

## The one rule

Read `PORT` from the environment and listen on every interface (`":"+port`, not
`"localhost:"+port`). velixir injects `PORT`; a hardcoded or loopback-only listener builds
fine and then fails its health check.

## Licence

MIT.
