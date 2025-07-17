FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN apk add --no-cache --update git make && make ci && make build

FROM alpine:latest

RUN apk add --no-cache --update ca-certificates tzdata

ENV TZ=Asia/Bangkok

WORKDIR /app

COPY --from=builder /app/main cmd/

RUN chmod +x cmd/main

EXPOSE 8080

CMD ["cmd/main"]