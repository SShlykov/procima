FROM golang:1.22 as builder
WORKDIR /app

COPY ./go.mod /app/go.mod

RUN go mod download
RUN go mod verify

COPY . .

RUN GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -o /app/main /app/cmd/procima

FROM alpine:3.14 as prod
WORKDIR /app

COPY --from=builder /app/main /app/main
COPY --from=builder /app/config /app/config

CMD [ "/app/main" ]