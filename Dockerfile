FROM golang:1.20-alpine AS build

RUN apk --update add ca-certificates upx && update-ca-certificates

WORKDIR /mini-health

RUN go env -w CGO_ENABLED="0"

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

RUN go build -ldflags="-s -w" -o /bin/mini-health . && upx --best --lzma /bin/mini-health

FROM scratch AS release

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /bin/mini-health /mini-health

ENTRYPOINT ["/mini-health"]
CMD ["https://example.org"]
