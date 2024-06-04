FROM golang:1.22.1-alpine3.18 AS builder

WORKDIR /app
COPY . .
RUN go mod vendor
RUN go build -o main ./cmd

FROM alpine:3.18

RUN apk --no-cache add ca-certificates tzdata chromium udev

WORKDIR /app

COPY --from=builder /app/main .
COPY .env .
COPY service_account.json .

ENV CHROME_BIN=/usr/bin/chromium-browser

EXPOSE 8888
CMD [ "/app/main" ]