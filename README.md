## Steps to Run

1. `minikube start`
2. `docker build --tag [username]/hello-go .`
3. `docker push [username]/hello-go:latest`
4. `kubectl create deployment demo --image [username]/hello-go`
5. `kubectl scale deployment demo --replicas=5`
6. `kubectl expose deployment demo --port 8888 --type=NodePort`
7. `minikube service demo`

```
curl -H "Connection: close" SERVICE_URL
```
