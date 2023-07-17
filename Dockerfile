FROM golang:1.20 AS builder

COPY . /src
WORKDIR /src

RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy && make build 

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app
COPY --from=builder /src/configs /app/configs

WORKDIR /app

CMD ["./cmd", "-conf", "configs"]

# docker run -itd --name call -p 8000:8000 callmysql