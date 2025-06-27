FROM golang:1.22.0-alpine as buildbase

RUN apk add git build-base ca-certificates

WORKDIR /go/src/github.com/rarimo/proof-verification-relayer

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go mod tidy && go mod vendor
RUN CGO_ENABLED=1 GO111MODULE=on GOOS=linux go build  -o /usr/local/bin/proof-verification-relayer /go/src/github.com/rarimo/proof-verification-relayer

FROM scratch
COPY --from=alpine:3.9 /bin/sh /bin/sh
COPY --from=alpine:3.9 /usr /usr
COPY --from=alpine:3.9 /lib /lib

COPY --from=buildbase /usr/local/bin/proof-verification-relayer /usr/local/bin/proof-verification-relayer
COPY --from=buildbase /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["proof-verification-relayer"]
