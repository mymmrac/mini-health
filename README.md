# Mini Health

[![CI](https://github.com/mymmrac/mini-health/actions/workflows/ci.yaml/badge.svg)](https://github.com/mymmrac/mini-health/actions/workflows/ci.yaml)
[![Docker Pulls](https://img.shields.io/docker/pulls/mymmrac/mini-health)](https://hub.docker.com/r/mymmrac/mini-health)

Simplest HTTP health check for scratch docker images

## Usage

In your Dockerfile:

```dockerfile
COPY --from=mymmrac/mini-health:latest /mini-health /mini-health
HEALTHCHECK CMD ["/mini-health", "https://example.org"]
```

Optionally copy ca-certificates if you don't have them already:

```dockerfile
COPY --from=mymmrac/mini-health:latest \
    /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
```

## Options

`mini-health [OPTIONS] URL`

| Option      | Description                                                                         |
|-------------|-------------------------------------------------------------------------------------|
| -c `int`    | Smallest HTTP status code that is considered as an error (default: 400 bad request) |
| -d `string` | Data passed as body of request                                                      |
| -e `string` | Name of environment variable used as base URL                                       |
| -m `string` | HTTP method used for request (default: GET)                                         |
| -q          | Quiet output (print only errors)                                                    |

### Example

```dockerfile
HEALTHCHECK CMD ["/mini-health", "-e", "BASE_URL", "-c", "500", "/health"]
```

In this example `GET` request to URL `$BASE_URL/health` with no body will be made,
HTTP status codes greater than `500` will be reported as errors.

## Install manually

```shell
go install github.com/mymmrac/mini-health@latest
```
