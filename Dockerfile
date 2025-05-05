FROM golang:latest AS builder

#工作目录
WORKDIR /gocloudisk

COPY go.mod ./
COPY go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o clouddisk && chmod +x clouddisk

# 使用更小的镜像运行构建产物
FROM alpine:latest

COPY --from=builder /gocloudisk/clouddisk /usr/bin/clouddisk

RUN chmod +x /usr/bin/clouddisk

ENTRYPOINT ["clouddisk"]