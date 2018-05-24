# GΛTEKEEPER

<img src="https://github.com/axoe/gatekeeper/blob/master/logo/logo.png" width="800">

----

Gatekeeper is an open source system for managing [secrets] in [AWS Secrets Manager]; providing basic mechanisms for creation, maintenance,
and rotation of secrets.


## To start developing Gatekeeper

This repository hosts all information about
building Gatekeeper from source, how to contribute code
and documentation.

If you want to build Gatekeeper right away there are two options:

##### You have a working [Go environment].

```
$ go get -d axeo/gatekeeper
$ cd $GOPATH/src/axeo/gatekeeper
$ gatekeeper
```

##### You have a working [Docker environment].

```
$ git clone https://github.com/axeo/gatekeeper
$ cd gatekeeper
$ docker build -t gatekeeper .
```

[secrets]: https://aws.amazon.com/secrets-manager/features/
[AWS Secrets Manager]: https://aws.amazon.com/secrets-manager/
