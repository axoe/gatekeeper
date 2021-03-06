# GΛTEKEEPER

[![Join the chat at https://gitter.im/gatekeeper-talk/Lobby](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gatekeeper-talk/?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

<img src=https://github.com/axoe/gatekeeper/blob/master/logo/logo-400px.png width="100">

----

Gatekeeper is an open source system for interacting with [secrets] in [AWS Secrets Manager]; providing basic mechanisms for creation,
and rotation of secrets.

## To start using Gatekeeper

### [Download Gatekeeper] - Linux, MacOS and Windows

Please download the proper package for your operating system and architecture. You can find the [SHA256 checksums] for Gatekeeper [here] and you can verify the checksums signature file which has been signed using AXOE's GPG key. 

### Overview

Gatekeeper is built on [cobra] and utilises a structure of commands, arguments & flags. It supports Fully POSIX-compliant flags (including short & long versions)

### Requirements

Gatekeeper uses the [AWS_PROFILE] environment variable

----

### Commands

 [**ls**] - List secrets in a specific region:

```
gatekeeper ls --region eu-west-2
```

 [**get**] - Get the value of a secret:

```
gatekeeper get --secret nameofsecret --region eu-west-2
```

[**add**] - Add a new secret:

```
gatekeeper add --name nameofsecret --region eu-west-2 --description "exampledescription" --value "{\"username\":\"foo\",\"password\":\"bar\"}"
```

[**update**] - Update a secret:

```
gatekeeper update --secret nameofsecret --region eu-west-2 --value "{\"username\":\"foo\",\"password\":\"foobar\"}"
```

[**del**] - Delete a secret:

```
gatekeeper del --secret nameofsecret --region eu-west-2
```

## To start developing Gatekeeper

This repository hosts all information about
building Gatekeeper from source, how to contribute code
and documentation.

If you want to build Gatekeeper right away there are two options:

##### You have a working [Go environment].

```
$ go get -d github.com/axoe/gatekeeper
$ cd $GOPATH/src/github.com/axoe/gatekeeper
$ gatekeeper
```

##### You have a working [Docker environment].

##### Build:
```
$ git clone https://github.com/axoe/gatekeeper
$ cd gatekeeper
$ docker build -t gatekeeper .
```

## Run Gatekeeper in Docker

By default the Docker image does not have an ENTRYPOINT set.
Uncomment the last line on the Dockerfile to launch Gatekeeper
when the container starts

##### Run:
```
$ docker run -it gatekeeper
```

##### Export [AWS ACCESS KEYS]:

Once you are inside the container you need to export your [AWS ACCESS KEYS]

```
/# export AWS_ACCESS_KEY_ID=AKIAXXXXXXXXXX
/# export AWS_SECRET_ACCESS_KEY=XxXxXxXXXX
/# gatekeeper ls --region eu-west-2
```

[Download Gatekeeper]: https://axoe.github.io/gatekeeper/
[here]: http://dl.tusker.io/gatekeeper/gatekeeper_0.1.0_SHA256SUMS
[SHA256 checksums]: http://dl.tusker.io/gatekeeper/gatekeeper_0.1.0_SHA256SUMS
[secrets]: https://aws.amazon.com/secrets-manager/features/
[AWS Secrets Manager]: https://aws.amazon.com/secrets-manager/
[cobra]: https://github.com/spf13/cobra
[AWS_PROFILE]: https://docs.aws.amazon.com/cli/latest/userguide/cli-multiple-profiles.html
[**ls**]: https://github.com/axoe/gatekeeper#commands
[**get**]: https://github.com/axoe/gatekeeper#commands
[**add**]: https://github.com/axoe/gatekeeper#commands
[**update**]: https://github.com/axoe/gatekeeper#commands
[**del**]: https://github.com/axoe/gatekeeper#commands
[AWS ACCESS KEYS]: https://docs.aws.amazon.com/cli/latest/userguide/cli-multiple-profiles.html
