# conf2kube

conf2kube creates Kubernetes secrets based on contents of configuration files.

This is not an official Google product.

## Usage

```
$ conf2kube -h
```
```
Usage of conf2kube:
  -f file
    	Path to configuration file.
  -n name
    	The name to use for the Kubernetes secret. Defaults to basename of configuration file.
```

### Create a new secret from a configuration file

```
conf2kube -f redis.conf | kubectl create -f -
```

### Update an existing secret

```
conf2kube -f redis.conf | kubectl apply -f -
```

### Print the contents of a secret

conf2kube can print the contents of a secret previously created by conf2kube.

```
kubectl get secrets redis.conf -o json | conf2kube
```

The secret must have a data element that matches the secret name. All secrets created by
conf2kube meet this requirement.

```
$ kubectl get secrets redis.conf -o yaml
```
```
apiVersion: v1
data:
  redis.conf: IyBSZWRpcyBj...
kind: Secret
metadata:
  creationTimestamp: 2015-11-30T04:13:14Z
  name: redis.conf
  namespace: default
  resourceVersion: "11049"
  selfLink: /api/v1/namespaces/default/secrets/redis.conf
  uid: ad06b82b-9718-11e5-a6e8-42010af00189
type: Opaque
```
