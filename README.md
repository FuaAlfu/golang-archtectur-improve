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

## storing passwords
never store passwords, instead store one-way encryption *"hash"* value of the password
for added security:
- hash on the client
- hash that again on the server
- hashing algorithms "bcrypt - current choice, scrypt - new choice"
[more info about scrybt](https://www.tarsnap.com/scrypt.html)


## Bearer Tokens & Hmac
* bearer tokens
- added to http spec with OAUTH2
- uses authorization header & keyword *"bearer"*
* to prevent faked bearer tokens, use cryptographic "signing"
- cryptographic signing is a way to provethat the value was created by certain person
-  HMAC
* Hmac
is a signing cryptographic algorithm "and that's all it is by itself" 

## jwt
json web token
{jtw standerd field}.{your fields}.Signature

## go module
```
go mod init folder-name or www.github.com/userName/repo-name
```

## latest packs
```
go mod tidy
```

## get version of module - analyzing a package to see if it's go module compatible
go list -m -version pkg-name
- example:
```
go list -m -versions github.com/dgrijalva/jwt-go
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