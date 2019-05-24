FROM golang:1.12 AS builder
WORKDIR /gorum

COPY go.mod ./
RUN go mod download

COPY cmd cmd
COPY pkg pkg

RUN CGO_ENABLED=0 go build -o /go/bin/server  -v ./cmd/server
RUN CGO_ENABLED=0 go build -o /go/bin/migrate -v ./cmd/migrate

FROM alpine:3.9 AS final
ENV HOST="0.0.0.0"
ENV PORT="80"

COPY --from=builder /go/bin/server  /bin/server
COPY --from=builder /go/bin/migrate /bin/migrate

COPY static    /usr/share/static
COPY templates /usr/share/templates
COPY config    /usr/share/config

COPY .docker/docker-entrypoint.sh /
ENTRYPOINT ["/docker-entrypoint.sh"]