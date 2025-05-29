### Useful commands
```bash
kubectl logs -f <pod-id> -n <context-name>

docker rm -f `docker ps -aq`

docker rmi -f `docker images -aq`

docker volume rm -f `docker volume ls -q`

docker logs <pod-id>

kubectl get pod -l service=<service-name> -n <context-name>

kubectl describe pod <pod-name> -n <context-name>

kubectl get pods -n <context-name> | grep <service-name>

aws ecr describe-images --repository-name <repository-name>

docker run -it --entrypoint=/bin/bash <url> -i  
```

### Kubernetes (Create own folder)
```bash
kubectl config use-context <name>

kubectl config set-context --current --namespace=<namespace>

kubectl get pods | grep <pod-name>
# OR
kubectl get pods -l group=<pod-name>

kubectl logs --follow <pod name>

# to see the previously crashed container's logs
kubectl logs --previous <pod name>

# If for any reason you need to run commands inside the pod you’ll need something like
kubectl exec -it <pod-id> -- bash
```
