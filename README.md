# JWT in Go

Just an example on how to use [JWT(JSON Web Token)](https://jwt.io/) in Go programming language.

> JSON Web Tokens are an open, industry standard RFC 7519 method for representing claims securely between two parties.

## Authentication

**Private key.** Signing
The private key is used to sign the generated token.

```shell
# Generate a private key
# '-out' is the flag to generate the file named 'private.rsa'
# 1024 is the bite size of the key
openssl genrsa -out private.rsa 1024
```

The command above generates a new file `private.rsa` with a new key:

```text
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQC3dJcFy3hXVly2Bd6feF0kb4t+dR+cGx7yO5m7/Aluqfem2YM4
b/ub68hLwLtAK+PjR5riFeSfzSr7Qo/b+9uexyWQmgsFWzKq2v7vMvPlt9cKTu6y
oM772I83J1Mz7UxIWeUn4EZ1kGjzX1oAaUCsid9r0HNWQApHEbzFQkSx5QIDAQAB
AoGAKdAH78ktN3K2lMIHuH79c6V8vKjDOJgx2HHfGypZmABjMoMW1lEnnqUBtMpL
o1edjBqw4WCuCAICDtxf33qq1kUbYmoP1Bp1/MILm/cHTzyTXomyDSp4ytCUcP8Z
02FYDn14AC1YBSZa/a+HxP5cN3dKPUH2/M3hLQdTucSIJYECQQDxslxhLe0a2ArA
IWU990ETbiWuLMhlDfIbrA2f1eiLojnDIfb2eBR8stoDglT5Oggr35jg+rTlVZZN
wLgB4ygFAkEAwk/jsnVM3LDy895g5iPec+5mtY2/TyxApMk4uNFIYRu9kMXNXbqG
MWA71TGExUUdi79Wx5F32McvRkJ5DJboYQJAJ56cRBXaMzdM89fFl9XLJhs9NAMF
SoNxt5WJmjTfbNxH6bPMnvRqL1LGKCMaOMyJgF+j2OVf9+QCwnYf87c/3QJBAIgH
1afBZqW5WtxO+hcjVUBjzZOcQCn1GOxD0mnQWZNiEDuhvju3sOowLLL7j69qy1xh
/yaHOBeSmg3dy1B89QECQFWhtnV/CFweo/CgPCjSHOxvWQMdG8LH75L2w8Mln+TS
ZHUcoyC6nxBZM6Vi7RvivU4pnHquyc7EiP4V4z0Gwxc=
-----END RSA PRIVATE KEY-----
```

**Public key.** Verify
The private key needs the public key to validate.

```shell
# Generate a public key using the private one
# '-in' indicates the private key file 'private.rsa'
# '-pubout >' saves to a file named 'public.rsa.pub'
openssl rsa -in private.rsa -pubout > public.rsa.pub
```

The command above generates a new file `public.rsa.pub` using as input the `private.rsa` file:

-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC3dJcFy3hXVly2Bd6feF0kb4t+
dR+cGx7yO5m7/Aluqfem2YM4b/ub68hLwLtAK+PjR5riFeSfzSr7Qo/b+9uexyWQ
mgsFWzKq2v7vMvPlt9cKTu6yoM772I83J1Mz7UxIWeUn4EZ1kGjzX1oAaUCsid9r
0HNWQApHEbzFQkSx5QIDAQAB
-----END PUBLIC KEY-----

## Request examples

`/login` endpoint to get a new JWT.

```shell
curl --location --request POST 'localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"avelino",
    "password":"password"
}'
```

`/validate` endpoint to verify and authorize the JWT.

```shell
curl --location --request POST 'localhost:8080/validate' \
--header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7Im5hbWUiOiJhdmVsaW5vIiwicm9sZSI6ImFkbWluIn0sImV4cCI6MTU5MDQ3NjE3NSwiaXNzIjoiaXNzdWVyIn0.LkQyNEr3wDRnAl2aOt2VQE95aMntQ0oWZqGRE_AwtcXq9YMxLNho_gvqjGiulRQMBpYmlaHzpN7t3ZsmYWrRSmUFMldR6V3xwo8bqPv4DTkIzDu19NpbFX_jjIUmlB8Y8g_clu8PC2_jqK3yZ3Jp2ujcZzkGnM75OtpWANCxmYc'
```
