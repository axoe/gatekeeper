FROM golang:latest as builder
WORKDIR /go/src/github.com/axeo/gatekeeper
RUN go get -d -v github.com/axeo/gatekeeper
COPY main.go  .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gatekeeper .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/axeo/gatekeeper/gatekeeper .
CMD ["./gatekeeper"]