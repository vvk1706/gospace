# Kubernetes Deployment Guide

This guide explains how to deploy GoSpace to Kubernetes using Rancher Desktop or any Kubernetes cluster.

## Overview

The application is deployed using Kubernetes manifests, which include:
- PostgreSQL StatefulSet with persistent storage ([`k8s-postgres.yaml`](k8s-postgres.yaml))
- Application Deployment with 2 replicas ([`k8s-deployment.yaml`](k8s-deployment.yaml))
- ConfigMap and Secret for database configuration
- NodePort service for external access
- Optional Ingress for domain-based routing
- Health checks (liveness and readiness probes)
- Resource limits and requests

## Prerequisites

- Rancher Desktop installed and running (or any Kubernetes cluster)
- kubectl configured to use your Kubernetes context
- Docker daemon accessible for building images

## Quick Deployment

### 1. Build Docker Image

First, build the Docker image in Rancher Desktop's Docker environment:

```bash
# Build the image
docker build -t gospace:latest .

# Verify the image
docker images | grep gospace
```

### 2. Deploy to Kubernetes

**Option A: Automated Deployment (Recommended)**

```bash
# Make script executable (first time only)
chmod +x deploy-k8s.sh

# Deploy everything (PostgreSQL + Application)
./deploy-k8s.sh

# The script will:
# - Create namespace
# - Deploy PostgreSQL with persistent storage
# - Deploy the application
# - Show access instructions
```

**Option B: Manual Deployment**

```bash
# Deploy PostgreSQL first
kubectl apply -f k8s-postgres.yaml

# Wait for PostgreSQL to be ready
kubectl wait --for=condition=ready pod -l app=postgres -n gospace --timeout=120s

# Deploy the application
kubectl apply -f k8s-deployment.yaml

# Verify deployment
kubectl get all -n gospace
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

### PostgreSQL Resources ([`k8s-postgres.yaml`](k8s-postgres.yaml))

#### ConfigMap
- **Name**: `postgres-config`
- **Namespace**: `gospace`
- **Data**:
  - `POSTGRES_DB`: gospace
  - `POSTGRES_USER`: postgres

#### Secret
- **Name**: `postgres-secret`
- **Namespace**: `gospace`
- **Data**:
  - `POSTGRES_PASSWORD`: Base64 encoded password

#### PersistentVolumeClaim
- **Name**: `postgres-pvc`
- **Namespace**: `gospace`
- **Storage**: 1Gi
- **Access Mode**: ReadWriteOnce
- **Storage Class**: local-path (Rancher Desktop default)

#### StatefulSet
- **Name**: `postgres`
- **Namespace**: `gospace`
- **Replicas**: 1
- **Image**: `postgres:15-alpine`
- **Container Port**: 5432
- **Volume Mount**: `/var/lib/postgresql/data`
- **Resources**:
  - Requests: 256Mi memory, 250m CPU
  - Limits: 512Mi memory, 500m CPU

#### Service
- **Name**: `postgres-service`
- **Namespace**: `gospace`
- **Type**: ClusterIP
- **Port**: 5432
- **Selector**: `app: postgres`

### Application Resources ([`k8s-deployment.yaml`](k8s-deployment.yaml))

### Namespace
- **Name**: `gospace`
- **Purpose**: Isolates all application resources
- **Labels**: `name: gospace`

### Deployment
- **Name**: `gospace`
- **Namespace**: `gospace`
- **Replicas**: 2 (for high availability)
- **Image**: `gospace:latest`
- **Image Pull Policy**: `IfNotPresent` (uses local images)
- **Container Port**: 8080
- **Environment Variables**:
  - `PORT`: "8080"
  - `DB_HOST`: postgres-service (from ConfigMap)
  - `DB_PORT`: "5432" (from ConfigMap)
  - `DB_USER`: postgres (from ConfigMap)
  - `DB_PASSWORD`: from Secret
  - `DB_NAME`: gospace (from ConfigMap)
  - `DB_SSLMODE`: disable (from ConfigMap)
- **Resources**:
  - Requests: 64Mi memory, 100m CPU
  - Limits: 128Mi memory, 200m CPU
- **Health Checks**:
  - **Liveness Probe**: HTTP GET on `/` (port 8080)
    - Initial delay: 10s
    - Period: 10s
    - Timeout: 5s
    - Failure threshold: 3
  - **Readiness Probe**: HTTP GET on `/` (port 8080)
    - Initial delay: 5s
    - Period: 5s
    - Timeout: 3s
    - Failure threshold: 3

### Service
- **Name**: `gospace-service`
- **Namespace**: `gospace`
- **Type**: NodePort
- **Selector**: `app: gospace`
- **Port**: 8080 (internal)
- **Target Port**: 8080
- **NodePort**: 30080 (external)
- **Access**: `http://localhost:30080`
- **Session Affinity**: None

### Ingress (Optional)
- **Name**: `gospace-ingress`
- **Namespace**: `gospace`
- **Host**: `gospace.local`
- **Ingress Class**: `nginx`
- **Annotations**: `nginx.ingress.kubernetes.io/rewrite-target: /`
- **Requires**: nginx ingress controller

## Useful Commands

### Check Deployment Status

```bash
# View all resources
kubectl get all -n gospace

# Check pod status
kubectl get pods -n gospace

# View pod logs
kubectl logs -n gospace -l app=gospace

# Follow logs
kubectl logs -n gospace -l app=gospace -f

# Describe deployment
kubectl describe deployment gospace -n gospace
```

### Scaling

```bash
# Scale to 3 replicas
kubectl scale deployment gospace -n gospace --replicas=3

# Verify scaling
kubectl get pods -n gospace
```

### Update Application

```bash
# Rebuild image
docker build -t gospace:latest .

# Restart deployment to use new image
kubectl rollout restart deployment gospace -n gospace

# Check rollout status
kubectl rollout status deployment gospace -n gospace

# View logs during update
kubectl logs -n gospace -l app=gospace -f
```

### Database Operations

```bash
# Connect to PostgreSQL
kubectl exec -it postgres-0 -n gospace -- psql -U postgres -d gospace

# Backup database
kubectl exec postgres-0 -n gospace -- pg_dump -U postgres gospace > backup.sql

# Restore database
kubectl exec -i postgres-0 -n gospace -- psql -U postgres gospace < backup.sql

# View database logs
kubectl logs postgres-0 -n gospace
```

### Access Pod Shell

```bash
# Get pod name
kubectl get pods -n gospace

# Access shell
kubectl exec -it <pod-name> -n gospace -- /bin/sh
```

### Port Forwarding (Alternative Access)

```bash
# Forward local port 8080 to service
kubectl port-forward -n gospace service/gospace-service 8080:8080

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
127.0.0.1 gospace.local
```

### 3. Access via Domain

```bash
open http://gospace.local
```

## Troubleshooting

### Pods Not Starting

```bash
# Check pod events
kubectl describe pod <pod-name> -n gospace

# Check logs
kubectl logs <pod-name> -n gospace
```

### Image Pull Issues

```bash
# Verify image exists
docker images | grep gospace

# Check imagePullPolicy in deployment
kubectl get deployment gospace -n gospace -o yaml | grep imagePullPolicy
```

### Service Not Accessible

```bash
# Check service endpoints
kubectl get endpoints -n gospace

# Verify service
kubectl describe service gospace-service -n gospace

# Test from within cluster
kubectl run -it --rm debug --image=alpine --restart=Never -n gospace -- sh
# Inside pod: wget -O- http://gospace-service:8080
```

### NodePort Not Working

```bash
# Verify NodePort is assigned
kubectl get service gospace-service -n gospace

# Check firewall rules (if applicable)
# For Rancher Desktop, usually no firewall issues on localhost
```

## Cleanup

Remove all resources:

```bash
# Delete all resources
kubectl delete -f k8s-deployment.yaml

# Verify deletion
kubectl get all -n gospace

# Delete namespace (if needed)
kubectl delete namespace gospace
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
The application uses PostgreSQL with persistent storage:
- PostgreSQL deployed as StatefulSet with PersistentVolumeClaim
- Data persists across pod restarts and redeployments
- 1Gi storage allocated by default (adjust in [`k8s-postgres.yaml`](k8s-postgres.yaml))
- For production, consider using a managed database service or increase storage size

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
  name: gospace
  namespace: gospace
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
                  - gospace
              topologyKey: kubernetes.io/hostname
      containers:
      - name: gospace
        image: gospace:latest
        resources:
          requests:
            memory: "128Mi"
            cpu: "200m"
          limits:
            memory: "256Mi"
            cpu: "500m"
```

## Additional Resources

- [`README.md`](README.md) - Main project documentation
- [`DEPLOYMENT.md`](DEPLOYMENT.md) - Comprehensive deployment guide
- [`API.md`](API.md) - API endpoint documentation
- [`QUICKSTART.md`](QUICKSTART.md) - Quick start guide
- [`docker-compose.yml`](docker-compose.yml) - Docker Compose configuration
- [`Dockerfile`](Dockerfile) - Docker image build configuration
- [`k8s-postgres.yaml`](k8s-postgres.yaml) - PostgreSQL Kubernetes manifest
- [`deploy-k8s.sh`](deploy-k8s.sh) - Automated Kubernetes deployment script

## Support

For issues or questions:
- Check pod logs: `kubectl logs -n gospace -l app=gospace`
- Review events: `kubectl get events -n gospace`
- Describe resources: `kubectl describe deployment gospace -n gospace`
- Consult Rancher Desktop documentation
- Review the project's GitHub issues