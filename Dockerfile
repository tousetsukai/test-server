FROM golang:alpine

WORKDIR /go
COPY . /go
RUN go build -o test-server
EXPOSE 8080
ENTRYPOINT ["./test-server"]
