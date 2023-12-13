# Docker
To develop locally, you need to install docker compose and run the below.
Once you build the images for the first time you no longer need to use --build.
But if there are any code changes I suggest that you do it again.
```
docker compose up --build
```

# Kubernetes
Make sure the tool `kubectl` or the TUI interface `k9s` is installed.

## Frontend development with live API 
```
$ kubectl port-forward svc/api-gateway-service 8000:8000

# Demo output:
# Forwarding from 127.0.0.1:8000 -> 8000
# Forwarding from [::1]:8000 -> 8000
```

## Backend development with live database
```
$ kubectl port-forward svc/postgresql 5432:5432

# Demo output:
# Forwarding from 127.0.0.1:5432 -> 5432
# Forwarding from [::1]:5432 -> 5432
```
