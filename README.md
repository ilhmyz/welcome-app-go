# welcome-app-go

### 1. Aplikasi welcome-app-go

Aplikasi ini dibuat dengan golang version 1.21+. untuk lokasi aplikasi berada di dir [./app/](./app/) aplikasi ini berjalan dengan route `/welcome` output 'Anonymous' dan `/welcome/{nama}` dan output Selamat datang {nama}

### 2. docker 
command untuk build image dengan tag `:testing`

``` sh
docker build -t welcome-app-go:testing .
```
karena Dockerfile berada di directory [`./docker/`](./docker/) , maka command untuk build
``` sh
docker build -f docker/Dockerfile -t welcome-app-go:testing .
```
untuk menajalankan container dengan image yang telah dibuat tadi, dengan port 5000 _(-p 8000:5000)_ tambahkan optiom `-d` agar container berjalan background
``` sh
docker run -p 8000:5000 welcome-app-go:testing
```
**docker-compose**

file docker compose berada di directory [`./docker/docker-compose.yaml`](./docker/docker-compose.yaml)

### 3. Github Actions
untuk github action berada di [.github/workflows](.github/workflows)
script tersebut akan berjalan dengan trigger manual karena menggunakan `workflow_dispatch` sebagai eventnya. script tersebut akan build docker image dan push ke dockerhub, deploy ke server dengan menggunakan docker/docker-compose (pastikan server sudah terinstal docker)

### 4. Kubernetes
untuk file script k8s semua berada di directory [`./kubernetes`](./kubernetes/)

1. ConfigMap &amp; Secret Script 
2. Deployment script
3. Service Script
4. Ingress Script

untuk https dengan letsencrypt terlebih dulu membuat issuer, karena bukan production maka menggunakan issuer staging

[`cert-staging-issuer.yaml`](./kubernetes/cert-staging-issuer.yaml)

```yaml
apiVersion: cert-manager.io/v1
kind: Issuer
metadata:
  name: letsencrypt-staging
spec:
  acme:
    server: https://acme-staging-v02.api.letsencrypt.org/directory
    email: user@example.com
    privateKeySecretRef:
      name: letsencrypt-staging
    solvers:
      - http01:
          ingress:
            ingressClassName: nginx
```
pada file ingress letsencrypt [`kubernetes/app-ingress-ssl.yaml`](./kubernetes/app-ingress-ssl.yaml) akan ditambahkan seperti ini
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: welcome-app-go
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
    cert-manager.io/issuer: "letsencrypt-staging"
spec:
  ingressClassName: nginx
  tls:
  - hosts:
    - local.host
    secretName: welcome-app-tls
  rules:
    - host: local.host
      http:
        paths:
          - pathType: Prefix
            path: /
            backend:
              service:
                name: welcome-app-go
                port:
                  number: 80
```


### 5. Infra Diagram
