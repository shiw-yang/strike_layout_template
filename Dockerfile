FROM golang:1.18.10 As builder

COPY . /src
WORKDIR /src

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN make build

FROM debian:stable-slim

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf/conf.yaml"]