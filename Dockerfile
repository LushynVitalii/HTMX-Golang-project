ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app .

# Final stage
FROM debian:bookworm

WORKDIR /app
COPY --from=builder /run-app /usr/local/bin/
COPY --from=builder /usr/src/app/assets /app/assets

EXPOSE 8080
CMD ["run-app"]