FROM golang:1.9.0-alpine3.6 AS build
RUN apk add --no-cache git
ADD . /go/src/github.com/axoe/gatekeeper/
RUN go get ./...
WORKDIR /go/src/github.com/axoe/gatekeeper/
RUN go install

FROM alpine:3.6
COPY --from=build /go/bin/gatekeeper /bin/gatekeeper
WORKDIR /go/src/github.com/axoe/gatekeeper/