# nginx Example

The following example will walk you through the creation of a Kubernetes secret named `nginx.conf`
based on the contents of the nginx.conf configuration file, as well as a secret containing the 
different ssl certificates and a pod running the nginx server to utilize them. 

## Create the nginx.conf secret

Create the nginx.conf Kubernetes secret from the nginx.conf configuration file using the
`conf2kube` and `kubectl create` commands.

```
$ conf2kube -f nginx.conf | kubectl create -f -
```

## Create the nginx-ssl-certs secrets

Create the nginx-ssl-certs Kubernetes secret from the certs using the
`conf2kube`, `kubectl create` and `kubectl patch` commands.

```
$ conf2kube -n nginx-ssl-certs -f example.com.crt | kubectl create -f -
$ kubectl patch secret nginx-ssl-certs -p `conf2kube -n nginx-ssl-certs -f example.com.csr`
$ kubectl patch secret nginx-ssl-certs -p `conf2kube -n nginx-ssl-certs -f example.com.key`
```

## Create the nginx pod

Create the nginx pod using the `kubectl create` command.

```
$ kubectl create -f nginx-pod.yaml
```

## Verify configuration settings

The nginx configuration settings can be verified online using telnet. First setup port forwarding
to the nginx pod on port 6443 using the `kubectl port-forward` command.

```
$ kubectl port-forward nginx 6443:443
```
```
Forwarding from 127.0.0.1:6443 -> 443
Forwarding from [::1]:6443 -> 443
```

Next, connect to the nginx server using the `curl` command. 

```
$ curl -k -H 'Host: example.com' https://127.0.0.1:6443
```
```
Handling connection for 6443
ok
```
