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
script tersebut akan berjalan dengan trigger manual karena menggunakan `workflow_dispatch` sebagai eventnya. script tersebut akan build docker image dan push ke dockerhub, deploy ke server dengan menggunakan docker/docker-compose

### 4. Kubernetes
untuk file script k8s semua berada di directory [`./kubernetes`](./kubernetes/)

1. ConfigMap &amp; Secret Script 
2. Deployment script
3. Service Script
4. Ingress Script

### 5. Infra Diagram
