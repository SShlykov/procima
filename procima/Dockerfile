FROM golang:1.22 as builder
WORKDIR /app

COPY ../procima/go.work.loc ./go.work

COPY ../procima/go.mod ./procima/go.mod
COPY ../go_pkg/logger/go.mod ./pkg_go/logger/go.mod
COPY ../go_pkg/metrics/go.mod ./pkg_go/metrics/go.mod

RUN go mod download
RUN go mod verify

COPY ../procima ./procima
COPY ../go_pkg/ ./pkg_go

RUN GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -o /app/main /app/procima/cmd/procima

FROM alpine:3.14
WORKDIR /app

COPY --from=builder /app/main /app/main
COPY ../procima/config /app/config

CMD [ "/app/main" ]