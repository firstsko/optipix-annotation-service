# 使用官方 Go 镜像进行构建
FROM golang:1.20 as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o optipix-socket ./cmd/main.go
FROM scratch
WORKDIR /app
COPY --from=builder /app/configs ./configs
COPY --from=builder /app/optipix-socket .
EXPOSE 8080

ENTRYPOINT ["/app/optipix-socket"]
