ARG IMAGE=scratch
ARG OS=linux
ARG ARCH=amd64

FROM golang:1.16.5-alpine3.12 as builder

WORKDIR /go/src/github.com/kinduff/csgo_exporter

RUN apk --no-cache --virtual .build-deps add git alpine-sdk

COPY . .

RUN GO111MODULE=on go mod vendor
RUN CGO_ENABLED=0 GOOS=$OS GOARCH=$ARCH go build -ldflags '-s -w' -o binary .

FROM $IMAGE

LABEL name="csgo_exporter"

WORKDIR /root/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/kinduff/csgo_exporter/binary csgo_exporter

CMD ["./csgo_exporter"]
