FROM golang:1.20.0
WORKDIR /go/src
ADD . /go/src
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN cd /go/src && go build -o main
CMD ["/go/src/main"]
