# conf2kube

conf2kube creates Kubernetes [secrets](http://kubernetes.io/v1.1/docs/user-guide/secrets.html)
based on contents of configuration files.

This is not an official Google product.

## Install

Download a [binary release](https://github.com/kelseyhightower/conf2kube/releases) or use the `go get` command.

```
$ go get github.com/kelseyhightower/conf2kube
```

## Usage

```
$ conf2kube -h
```
```
Usage of conf2kube:
  -f file
        Path to configuration file. Defaults to stdin. (default "-")
  -n name
        The name to use for the Kubernetes secret. Defaults to basename of configuration file.
  -x    Extract config from incoming JSON formated secret and print to stdout.
```

### Create a new secret from stdin

```
$ cat redis.conf | conf2kube -n redis.conf -f - | kubectl create -f -
```

> The `-f` flag is optional when creating secrets from stdin.

### Create a new secret from a configuration file

```
$ conf2kube -f redis.conf | kubectl create -f -
```

Use the `-n` flag to create a secret with a specific name.

```
$ conf2kube -f redis.conf -n redis-master.conf | kubectl create -f -
```

### Update an existing secret

Use the `kubectl apply` command to update an existing secret.

```
$ conf2kube -f redis.conf | kubectl apply -f -
```

### Extract the contents of a secret

Use the `-x` flag to extract JSON formated secrets from stdin.

```
$ kubectl get secrets redis.conf -o json | conf2kube -x
```

Use the `-n` flag to extract a secret by a specific data name. Useful when the data name does
not match the secret name.

```
$ kubectl get secrets redis.conf -o json | conf2kube -x -n redis-master.conf
```
