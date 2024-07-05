ARG goVersion=1.16-alpine

FROM golang:${goVersion} as builder
ARG gitCommit=""
ARG buildStamp=""

ENV GO111MODULE=on TZ=Asia/Shanghai

WORKDIR /app

ADD . .

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && \
    apk --no-cache add bash tzdata

RUN go env -w  GOPROXY=https://goproxy.cn,direct \
    && go build -ldflags "-s -w -X 'main.gitCommit=${gitCommit}' -X 'main.buildStamp=${buildStamp}'" -o go-server cmd/server/main.go

FROM alpine

LABEL author=renjie email=penitence.rj@gmail.com

ENV GIN_MODE=release TZ=Asia/Shanghai

RUN mkdir -p /app

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && \
    apk --no-cache add bash tzdata

COPY --from=builder /app/go-server /app
COPY --from=builder /app/app.yml /app
COPY --from=builder /app/sql /app/sql

WORKDIR app

EXPOSE 8080

CMD ["./go-server"]