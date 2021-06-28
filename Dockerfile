FROM golang:1.16.0-alpine AS builder
WORKDIR /build
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
COPY . .
RUN go mod init exponea.com && go mod tidy && go build

FROM scratch
EXPOSE 8080
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/exponea.com /
ENTRYPOINT ["/exponea.com"]
