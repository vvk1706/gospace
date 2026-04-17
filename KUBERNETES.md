# Kubernetes Deployment Guide

This guide explains how to deploy the Gin Web Application to Kubernetes using Rancher Desktop.

## Prerequisites

- Rancher Desktop installed and running
- kubectl configured to use Rancher Desktop context
- Docker daemon accessible from Rancher Desktop

## Quick Deployment

### 1. Build Docker Image

First, build the Docker image in Rancher Desktop's Docker environment:

```bash
# Build the image
docker build -t gin-webapp:latest .

# Verify the image
docker images | grep gin-webapp
```

### 2. Deploy to Kubernetes

Apply the Kubernetes manifests:

```bash
# Deploy all resources
kubectl apply -f k8s-deployment.yaml

# Verify deployment
kubectl get all -n gin-webapp
```

### 3. Access the Application

The application is exposed via NodePort on port 30080:

```bash
# Get the node IP (usually localhost for Rancher Desktop)
kubectl get nodes -o wide

# Access the application
open http://localhost:30080
```

## Kubernetes Resources

The deployment includes:

### Namespace
- **Name**: `gin-webapp`
- Isolates all application resources

### Deployment
- **Name**: `gin-webapp`
- **Replicas**: 2 (for high availability)
- **Image**: `gin-webapp:latest`
- **Container Port**: 8080
- **Resources**:
  - Requests: 64Mi memory, 100m CPU
  - Limits: 128Mi memory, 200m CPU
- **Health Checks**:
  - Liveness probe on `/` endpoint
  - Readiness probe on `/` endpoint

### Service
- **Name**: `gin-webapp-service`
- **Type**: NodePort
- **Port**: 8080 (internal)
- **NodePort**: 30080 (external)
- **Access**: `http://localhost:30080`

### Ingress (Optional)
- **Name**: `gin-webapp-ingress`
- **Host**: `gin-webapp.local`
- Requires nginx ingress controller

## Useful Commands

### Check Deployment Status

```bash
# View all resources
kubectl get all -n gin-webapp

# Check pod status
kubectl get pods -n gin-webapp

# View pod logs
kubectl logs -n gin-webapp -l app=gin-webapp

# Follow logs
kubectl logs -n gin-webapp -l app=gin-webapp -f

# Describe deployment
kubectl describe deployment gin-webapp -n gin-webapp
```

### Scaling

```bash
# Scale to 3 replicas
kubectl scale deployment gin-webapp -n gin-webapp --replicas=3

# Verify scaling
kubectl get pods -n gin-webapp
```

### Update Application

```bash
# Rebuild image
docker build -t gin-webapp:latest .

# Restart deployment to use new image
kubectl rollout restart deployment gin-webapp -n gin-webapp

# Check rollout status
kubectl rollout status deployment gin-webapp -n gin-webapp
```

### Access Pod Shell

```bash
# Get pod name
kubectl get pods -n gin-webapp

# Access shell
kubectl exec -it <pod-name> -n gin-webapp -- /bin/sh
```

### Port Forwarding (Alternative Access)

```bash
# Forward local port 8080 to service
kubectl port-forward -n gin-webapp service/gin-webapp-service 8080:8080

# Access at http://localhost:8080
```

## Using Ingress (Optional)

If you want to use domain-based access:

### 1. Install Nginx Ingress Controller

```bash
# Install nginx ingress
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.1/deploy/static/provider/cloud/deploy.yaml
```

### 2. Configure Local DNS

Add to `/etc/hosts`:

```
127.0.0.1 gin-webapp.local
```

### 3. Access via Domain

```bash
open http://gin-webapp.local
```

## Troubleshooting

### Pods Not Starting

```bash
# Check pod events
kubectl describe pod <pod-name> -n gin-webapp

# Check logs
kubectl logs <pod-name> -n gin-webapp
```

### Image Pull Issues

```bash
# Verify image exists
docker images | grep gin-webapp

# Check imagePullPolicy in deployment
kubectl get deployment gin-webapp -n gin-webapp -o yaml | grep imagePullPolicy
```

### Service Not Accessible

```bash
# Check service endpoints
kubectl get endpoints -n gin-webapp

# Verify service
kubectl describe service gin-webapp-service -n gin-webapp

# Test from within cluster
kubectl run -it --rm debug --image=alpine --restart=Never -n gin-webapp -- sh
# Inside pod: wget -O- http://gin-webapp-service:8080
```

### NodePort Not Working

```bash
# Verify NodePort is assigned
kubectl get service gin-webapp-service -n gin-webapp

# Check firewall rules (if applicable)
# For Rancher Desktop, usually no firewall issues on localhost
```

## Cleanup

Remove all resources:

```bash
# Delete all resources
kubectl delete -f k8s-deployment.yaml

# Verify deletion
kubectl get all -n gin-webapp

# Delete namespace (if needed)
kubectl delete namespace gin-webapp
```

## Production Considerations

### High Availability
- Increase replicas: `replicas: 3` or more
- Use pod anti-affinity to spread across nodes
- Configure horizontal pod autoscaling (HPA)

### Resource Management
```yaml
resources:
  requests:
    memory: "128Mi"
    cpu: "200m"
  limits:
    memory: "256Mi"
    cpu: "500m"
```

### Persistent Storage
Since the app uses in-memory storage, data is lost on pod restart. For production:
- Consider adding a real database
- Use StatefulSet for stateful workloads
- Add PersistentVolumeClaims for data persistence

### Security
- Use non-root user in container
- Enable Pod Security Standards
- Add network policies
- Use secrets for sensitive data

### Monitoring
- Add Prometheus metrics endpoint
- Configure ServiceMonitor for Prometheus
- Set up Grafana dashboards
- Enable logging aggregation

## Example: Production-Ready Configuration

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: gin-webapp
  namespace: gin-webapp
spec:
  replicas: 3
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0
  template:
    spec:
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - weight: 100
            podAffinityTerm:
              labelSelector:
                matchExpressions:
                - key: app
                  operator: In
                  values:
                  - gin-webapp
              topologyKey: kubernetes.io/hostname
      containers:
      - name: gin-webapp
        image: gin-webapp:latest
        resources:
          requests:
            memory: "128Mi"
            cpu: "200m"
          limits:
            memory: "256Mi"
            cpu: "500m"
```

## Support

For issues or questions:
- Check pod logs: `kubectl logs -n gin-webapp -l app=gin-webapp`
- Review events: `kubectl get events -n gin-webapp`
- Consult Rancher Desktop documentation