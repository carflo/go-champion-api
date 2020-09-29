FROM golang:1.15-alpine as build

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080/TCP

CMD ["app"]

