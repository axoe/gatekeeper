FROM golang:1.10-alpine3.7 AS build
RUN apk add --no-cache git
ADD . /go/src/github.com/axoe/gatekeeper/
RUN go get ./...
WORKDIR /go/src/github.com/axoe/gatekeeper/
RUN go install

FROM alpine:3.7
COPY --from=build /go/bin/gatekeeper /bin/gatekeeper
WORKDIR /bin
#ENTRYPOINT ["/bin/gatekeeper"]