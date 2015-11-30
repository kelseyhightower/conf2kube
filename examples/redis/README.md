# Redis Example

The following example will walk you through the creation of a Kubernetes secret named `redis.conf`
based on the contents of the redis.conf configuration file, and a pod running the redis server to
utilize it. 

## Create the redis.conf secret

Create the redis.conf Kubernetes secret from the redis.conf configuration file using the
`conf2kube` and `kubectl create` commands.

```
$ conf2kube -f redis.conf | kubectl create -f -
```

## Create the redis pod

Create the redis pod using the `kubectl create` command.

```
$ kubectl create -f redis-pod.yaml
```

## Verify configuration settings

The redis configuration settings can be verified online using telnet. First setup port forwarding
to the redis pod on port 6379 using the `kubectl port-forward` command.

```
$ kubectl port-forward redis 6379
```
```
Forwarding from 127.0.0.1:6379 -> 6379
Forwarding from [::1]:6379 -> 6379
```

Next, connect to the redis server using the `telnet` command. Type `CONFIG GET loglevel` at the
prompt to retrieve the loglevel configuration setting.

```
$ telnet localhost 6379
```
```
Trying ::1...
Connected to localhost.
Escape character is '^]'.
CONFIG GET loglevel
*2
$8
loglevel
$5
notice
```

## Updating the redis.conf secret

First edit the redis.conf configuration file, then use the `kubectl apply` command to post the results.

```
$ conf2kube -f redis.conf | kubectl apply -f -
```

Next, recreate the redis pod using the `kubectl` command.

```
$ kubectl delete pods redis
```
```
$ kubectl create -f redis-pod.yaml
```

At this point the redis pod should be using the updated redis.conf Kubernetes secret.
