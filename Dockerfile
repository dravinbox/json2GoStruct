FROM golang:1.11.0 as builder

WORKDIR /go/src/json2GoStruct/

RUN go get -d -v github.com/kataras/iris

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o webapp .



FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/json2GoStruct .

EXPOSE 8080

CMD ["./webapp"]