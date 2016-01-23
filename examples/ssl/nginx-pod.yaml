apiVersion: v1
kind: Pod
metadata:
  name: nginx
  labels:
    name: nginx
spec:
  containers:
  - name: nginx
    image: nginx
    ports:
      - containerPort: 443
    volumeMounts:
      - mountPath: "/etc/nginx/"
        name: "nginx-conf"
      - mountPath: "/usr/local/etc/nginx/ssl"
        name: "ssl-certs"
  volumes:
    - name: "nginx-conf"
      secret:
        secretName: "nginx.conf"
    - name: "ssl-certs"
      secret:
        secretName: "nginx-ssl-certs"