FROM golang:1.18

WORKDIR /usr/src/app

COPY . .
RUN go build

COPY config.yml.example config.yml

RUN chmod +x daemon.sh

CMD ["./daemon.sh"]
