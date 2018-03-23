FROM golang:1.8

ENV https_proxy http://10.144.1.10:8080

WORKDIR /go/src/app
COPY . .

RUN go get -d -v "gopkg.in/alecthomas/kingpin.v2"
RUN go get -d -v "github.com/Shopify/sarama"
