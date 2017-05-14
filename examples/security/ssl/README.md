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

Start a secured etcd cluster with 3-nodes through `goreman`

```
$ goreman start
```
