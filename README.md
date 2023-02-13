# Mini Health

[![CI](https://github.com/mymmrac/mini-health/actions/workflows/ci.yaml/badge.svg)](https://github.com/mymmrac/mini-health/actions/workflows/ci.yaml)

Simplest HTTP health check for scratch docker images 

## Usage

In your Dockerfile:

```dockerfile
COPY --from=mymmrac/mimi-health:latest /mini-health /mini-health
HEALTHCHECK CMD /mini-health https://example.org
```

Optionally copy ca-certificates if you don't have them already:

```dockerfile
COPY --from=mymmrac/mimi-health:latest \
    /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
```

## Install manually

```shell
go install github.com/mymmrac/mini-health@latest
```
