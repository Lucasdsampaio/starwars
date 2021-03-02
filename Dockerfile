FROM golang:1.15.4-alpine3.12 AS builder

WORKDIR /starwars

COPY ./src .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build --ldflags='-w -s -extldflags "-static"' -v -a -o /go/bin/starwars .

FROM alpine:3.12

ENV GIN_MODE=release

COPY --from=builder /go/bin/starwars /go/bin/starwars

EXPOSE 8080

ENTRYPOINT ["/go/bin/starwars"]