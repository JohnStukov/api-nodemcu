FROM golang:1.16.0-alpine3.13

WORKDIR /go/src/github.com/Ignis-Divine/api-nodemcu
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["api-nodemcu"]