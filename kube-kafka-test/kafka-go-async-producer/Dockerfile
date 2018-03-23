FROM kafka-go:v1

ENV https_proxy http://10.144.1.10:8080

WORKDIR /go/src/app
COPY . .

#RUN go get -d -v ./...
#RUN go install -v ./...

CMD [ "go", "run", "./async-producer.go" ]
