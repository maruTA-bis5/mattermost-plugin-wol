FROM golang:1.12

VOLUME ["/go", "/repo"]
RUN mkdir /.cache && chmod 777 /.cache && chmod 777 /go
WORKDIR /repo
