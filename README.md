---
stack: GO
lang: all
---

## web authentication with golang
all about Encryption, jwt, Oauth, HMAC and more..

## online crul builder
[click here](https://tools.w3cub.com/curl-builder)

## basic authentication
basic authentication part of the specification of http
send username / password with every request.
uses authorization header & keyword *"basic"*
- put *"uesrname:password"* together.
- converts them to *base64*.
- *basic64* put generic binary data into form
- *base64 is reversable*, never use with http, only https
- use basic athentication to login.
for example:
```
crul -u user:pass -v google.com
```

## go module
```
go mod init folder-name or www.github.com/userName/repo-name
```

## latest packs
```
go mod tidy
```

## Got a problem with lunch? GOPATH should be set to
```
export GOPATH=$GOROOT
unset GOROOT
```

##  GO111MODULE set "on" or "off"
```
go env -w GO111MODULE=off
```