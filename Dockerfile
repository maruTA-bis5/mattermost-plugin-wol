FROM golang:1.24

VOLUME ["/go", "/repo"]
RUN mkdir /.cache && chmod 777 /.cache && chmod 777 /go
WORKDIR /repo
