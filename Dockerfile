FROM golang:latest as builder
WORKDIR /go/src/github.com/axoe/gatekeeper
RUN go get -d -v github.com/axoe/gatekeeper
COPY main.go  .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gatekeeper .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/axoe/gatekeeper/gatekeeper .
CMD ["./gatekeeper"]