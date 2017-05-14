ssl example
----------------

Getting the correct tools together is a black art, and can quickly become a debugging nightmare if it's not done correctly.

This tutorial works you through a working example to secure your etcd cluster.

# Prerequisite

Make sure following packages are installed properly on your hosts:

* [Go 1.8+ (if you want to install etcd and cfssl from source)](https://golang.org/dl/)
* [goreman](https://github.com/mattn/goreman)
* [cfssl](https://github.com/cloudflare/cfssl)

# Generate CA and Certificates

Configure the example in `config`

|Variable|Default Value|Description|
|--------|-------------|-----------|
|DEFAULT_ADDR|127.0.0.1|hosts separated by comma, specifying default authorized hosts across CA, server, client and peer certificates|
|SERVER_NAME|server|CN field in server certificate|
|SERVER_ADDR|$DEFAULT_ADDR|set hosts for server certificate|
|CLIENT_NAME|client|CN field in client certificate|
|CLIENT_ADDR|$DEFAULT_ADDR|set hosts for client certificate|
|MEMBERS|infra1 infra2 infra3|names specifying peers in the etcd cluster|

Generate CA and certificates with `gen.sh`

```
$ sh gen.sh
```

`gen.sh` creates a self-signed CA, along with separated certificates for server,
client and configured members(or peers),
and generated certificates are signed by the self-signed CA.

Start a secured etcd cluster with 3-nodes through `goreman`

```
$ goreman start
```

Perform a secured operation via `etcdctl`

```
$ etcdctl --endpoints=https://127.0.0.1:12379,https://127.0.0.1:22379,https://127.0.0.1:32379 --cert-file=client.pem --ca-file=ca.pem --key-file=client-key.pem member list

3fb0bb8f51408909: name=infra3 peerURLs=https://127.0.0.1:32380 clientURLs=https://127.0.0.1:32379 isLeader=false
d05e7521f6de6bab: name=infra1 peerURLs=https://127.0.0.1:12380 clientURLs=https://127.0.0.1:12379 isLeader=true
f0c6bf0f0690ad50: name=infra2 peerURLs=https://127.0.0.1:22380 clientURLs=https://127.0.0.1:22379 isLeader=false
```

# Reference

[Introducing CFSSL](https://blog.cloudflare.com/introducing-cfssl/)

[Self-signed Certificates](https://en.wikipedia.org/wiki/Self-signed_certificate)

[Public Key Architecture](https://en.wikipedia.org/wiki/Public_key_infrastructure)
