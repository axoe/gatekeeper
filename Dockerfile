FROM golang:1.10-alpine3.7 AS build
RUN apk update && apk add git ca-certificates && rm -rf /var/cache/apk/*
ADD . /go/src/github.com/axoe/gatekeeper/
RUN go get ./...
WORKDIR /go/src/github.com/axoe/gatekeeper/
RUN go install

FROM alpine:3.7
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
COPY --from=build /go/bin/gatekeeper $GOPATH/bin

# Comment out the below line if you wish to use Gatekeeper via the shell in CI
#ENTRYPOINT ["gatekeeper"]