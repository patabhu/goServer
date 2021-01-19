FROM golang:alpine AS gobuilder
RUN apk add git
ADD . /go/src/goServer
WORKDIR /go/src/goServer
RUN go build -o main
FROM alpine:latest 
WORKDIR /root/
COPY --from=gobuilder /go/src/goServer/main .
CMD ["./main"]