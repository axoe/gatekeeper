FROM golang:1.10-alpine3.7 AS golang
ADD . /go/src/github.com/axoe/gatekeeper/
WORKDIR /go/src/github.com/axoe/gatekeeper/
RUN apk update \
    && apk add git ca-certificates \
    && rm -rf /var/cache/apk/* \
    && CGO_ENABLED=0 go get -d -v ./... && go install -v -ldflags '-extldflags "-static"'

FROM alpine:3.7
COPY --from=golang /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
COPY --from=golang /go/bin/gatekeeper $GOPATH/bin

# Uncomment and populate to add AWS Keys to the ENV on
#ENV AWS_ACCESS_KEY_ID=AKIAXXXX
#ENV AWS_SECRET_ACCESS_KEY=XxXxXXX

# Uncomment the below line if you want gatekeeper to start on launch
#ENTRYPOINT ["gatekeeper"]