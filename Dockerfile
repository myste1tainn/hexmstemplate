FROM golang:1.18.0 as builder

WORKDIR /app

COPY go.* .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build \
    -ldflags "-X main.githash=$(git rev-parse HEAD) -X main.buildstamp=$(date +%Y%m%d.%H%M%S)" \
    -o goapp main.go


##########################################

FROM alpine:latest

RUN apk update && apk add --no-cache \
        ca-certificates \
        tzdata

ENV TZ=Asia/Bangkok


WORKDIR /app

COPY ./configs ./configs
COPY --from=builder /app/goapp ./goapp

EXPOSE 1323

CMD ["./goapp"]

