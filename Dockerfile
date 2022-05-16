FROM golang:1.17.5-alpine3.15 AS builder
WORKDIR /app
ENV GOPROXY https://goproxy.cn
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build  -ldflags "-s -w" -a -o gitlab-dingtalk .
RUN apk update && apk add upx && upx -9 -o dingtalk gitlab-dingtalk


FROM alpine:3.15
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /app/dingtalk /app/gitlab-dingtalk
COPY --from=builder /app/tpls /app/tpls
COPY --from=builder /app/application.yml /app/application.yml
WORKDIR /app
ENTRYPOINT ["./gitlab-dingtalk"]
