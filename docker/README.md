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
