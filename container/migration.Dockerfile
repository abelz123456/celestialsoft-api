FROM golang:1.20.4 AS build
WORKDIR /go/src/boilerplate

COPY . .
RUN rm -rf ./main.go && \
    mv ./cmd/db_migration.go ./main.go

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

FROM alpine:latest as release
RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=build /go/src/boilerplate .
RUN rm -rf ./main.go

RUN apk -U upgrade \
    && apk add --no-cache dumb-init ca-certificates \
    && chmod +x /app/app

CMD ["./app", "-prod"]
ENTRYPOINT ["/usr/bin/dumb-init", "--"]