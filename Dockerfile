FROM golang:latest as builder
WORKDIR /go/src/github.com/mrmoneyc/web-demo-golang/
COPY src/main.go /go/src/github.com/mrmoneyc/web-demo-golang/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM nvidia/cuda:8.0-cudnn5-devel-ubuntu14.04
WORKDIR /opt/web-demo-golang/
COPY --from=builder /go/src/github.com/mrmoneyc/web-demo-golang/server .
CMD ["./server"]
