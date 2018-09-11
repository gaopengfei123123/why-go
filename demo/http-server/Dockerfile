# 编译容器
FROM golang:1.9-alpine AS builder
WORKDIR /go/src/http-demo/
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# 运行容器
FROM alpine:latest AS runner
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/http-demo/app .
EXPOSE 9090
CMD [ "./app" ]