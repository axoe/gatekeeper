# GΛTEKEEPER

[![Join the chat at https://gitter.im/gatekeeper-talk/Lobby](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gatekeeper-talk/?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

<img src="https://github.com/axoe/gatekeeper/blob/master/logo/logo.png" width="100">

----

Gatekeeper is an open source system for managing [secrets] in [AWS Secrets Manager]; providing basic mechanisms for creation, maintenance,
and rotation of secrets.

## To start using Gatekeeper

### Overview

Gatekeeper is built on [cobra] and utilises a structure of commands, arguments & flags. It supports Fully POSIX-compliant flags (including short & long versions)

### Requirements

Gatekeeper uses the [AWS_PROFILE] environment variable

----

### Commands

List secrets in a specific region:

```
gatekeeper ls --region eu-west-2
```

Get the value of a secret:

```
gatekeeper get --secret nameofsecret --region eu-west-2
```

Add a new secret:

```
gatekeeper add --name nameofsecret --region eu-west-2 --description "exampledescription" --value "{\"username\":\"foo\",\"password\":\"bar\"}"
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

```
$ git clone https://github.com/axoe/gatekeeper
$ cd gatekeeper
$ docker build -t gatekeeper .
```

[secrets]: https://aws.amazon.com/secrets-manager/features/
[AWS Secrets Manager]: https://aws.amazon.com/secrets-manager/
[cobra]: https://github.com/spf13/cobra
[AWS_PROFILE]: https://docs.aws.amazon.com/cli/latest/userguide/cli-multiple-profiles.html