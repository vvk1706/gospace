# GoSpace Deployment Guide

This guide covers all deployment options for the GoSpace application.

## Table of Contents

- [Docker Deployment](#docker-deployment)
- [Kubernetes Deployment](#kubernetes-deployment)
- [Manual Deployment](#manual-deployment)
- [Environment Variables](#environment-variables)
- [Troubleshooting](#troubleshooting)

## Docker Deployment

### Prerequisites

- Docker 20.10+
- Docker Compose 2.0+

### Quick Deploy

Use the automated deployment script:

```bash
./deploy-docker.sh
```

### Manual Deploy

1. **Create environment file**

```bash
cp .env.example .env
# Edit .env with your configuration
```

2. **Build and start services**

```bash
docker build -t gospace:latest .
docker-compose up -d
```

3. **Verify deployment**

```bash
docker-compose ps
docker-compose logs -f app
```

4. **Access application**

- Application: http://localhost:8080
- PostgreSQL: localhost:5432

### Docker Compose Commands

```bash
# View logs
docker-compose logs -f

# Restart services
docker-compose restart

# Stop services
docker-compose down

# Remove all data
docker-compose down -v

# Rebuild and restart
docker-compose up -d --build
```

## Kubernetes Deployment

### Prerequisites

- Kubernetes cluster (minikube, kind, or cloud provider)
- kubectl configured
- Docker for building images

### Quick Deploy

Use the automated deployment script:

```bash
./deploy-k8s.sh
```

### Manual Deploy

1. **Build Docker image**

```bash
docker build -t gospace:latest .
```

2. **Deploy PostgreSQL**

```bash
kubectl apply -f k8s-postgres.yaml
```

3. **Wait for PostgreSQL to be ready**

```bash
kubectl wait --for=condition=ready pod -l app=postgres -n gospace --timeout=120s
```

4. **Deploy application**

```bash
kubectl apply -f k8s-deployment.yaml
```

5. **Verify deployment**

```bash
kubectl get all -n gospace
kubectl logs -f deployment/gospace -n gospace
```

6. **Access application**

- NodePort: http://localhost:30080
- Port Forward: `kubectl port-forward -n gospace svc/gospace-service 8080:8080`
- Ingress: http://gospace.local (requires ingress controller)

### Kubernetes Commands

```bash
# View all resources
kubectl get all -n gospace

# View pods
kubectl get pods -n gospace

# View logs
kubectl logs -f deployment/gospace -n gospace

# Describe pod
kubectl describe pod <pod-name> -n gospace

# Execute command in pod
kubectl exec -it <pod-name> -n gospace -- /bin/sh

# Scale deployment
kubectl scale deployment gospace -n gospace --replicas=3

# Delete deployment
kubectl delete namespace gospace
```

### Kubernetes Architecture

```
┌─────────────────────────────────────────┐
│           Namespace: gospace            │
├─────────────────────────────────────────┤
│                                         │
│  ┌──────────────┐    ┌──────────────┐  │
│  │   GoSpace    │    │  PostgreSQL  │  │
│  │  Deployment  │───▶│  Deployment  │  │
│  │  (2 replicas)│    │  (1 replica) │  │
│  └──────────────┘    └──────────────┘  │
│         │                    │          │
│         ▼                    ▼          │
│  ┌──────────────┐    ┌──────────────┐  │
│  │   Service    │    │   Service    │  │
│  │  (NodePort)  │    │ (ClusterIP)  │  │
│  └──────────────┘    └──────────────┘  │
│         │                    │          │
│         ▼                    ▼          │
│  ┌──────────────┐    ┌──────────────┐  │
│  │   Ingress    │    │     PVC      │  │
│  │  (Optional)  │    │   (1Gi)      │  │
│  └──────────────┘    └──────────────┘  │
│                                         │
└─────────────────────────────────────────┘
```

## Manual Deployment

### Prerequisites

- Go 1.21+
- PostgreSQL 15+

### Steps

1. **Install PostgreSQL**

```bash
# macOS
brew install postgresql@15

# Ubuntu/Debian
sudo apt-get install postgresql-15

# Start PostgreSQL
brew services start postgresql@15  # macOS
sudo systemctl start postgresql    # Linux
```

2. **Create database**

```bash
psql postgres
CREATE DATABASE gospace;
CREATE USER gospace_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE gospace TO gospace_user;
\q
```

3. **Configure environment**

```bash
cp .env.example .env
# Edit .env with your PostgreSQL credentials
```

4. **Build and run**

```bash
go mod download
go build -o gospace main.go
./gospace
```

## Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| PORT | Application port | 8080 | No |
| DB_HOST | PostgreSQL host | localhost | Yes |
| DB_PORT | PostgreSQL port | 5432 | Yes |
| DB_USER | PostgreSQL user | postgres | Yes |
| DB_PASSWORD | PostgreSQL password | - | Yes |
| DB_NAME | Database name | gospace | Yes |
| DB_SSLMODE | SSL mode | disable | No |

### Example .env file

```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_secure_password
DB_NAME=gospace
DB_SSLMODE=disable
```

## Troubleshooting

### Docker Issues

**Problem**: Container fails to start

```bash
# Check logs
docker-compose logs app

# Check if port is already in use
lsof -i :8080

# Rebuild without cache
docker-compose build --no-cache
docker-compose up -d
```

**Problem**: Database connection fails

```bash
# Check PostgreSQL is running
docker-compose ps postgres

# Check database logs
docker-compose logs postgres

# Restart PostgreSQL
docker-compose restart postgres
```

### Kubernetes Issues

**Problem**: Pods not starting

```bash
# Check pod status
kubectl get pods -n gospace

# Describe pod for events
kubectl describe pod <pod-name> -n gospace

# Check logs
kubectl logs <pod-name> -n gospace
```

**Problem**: Image pull errors

```bash
# For local images, ensure imagePullPolicy is IfNotPresent
# Load image into cluster (for minikube/kind)
minikube image load gospace:latest
kind load docker-image gospace:latest
```

**Problem**: Database connection fails

```bash
# Check PostgreSQL pod
kubectl get pods -l app=postgres -n gospace

# Check PostgreSQL logs
kubectl logs -l app=postgres -n gospace

# Verify service
kubectl get svc postgres-service -n gospace
```

### Application Issues

**Problem**: Application crashes on startup

```bash
# Check environment variables
docker-compose exec app env | grep DB_

# Test database connection
docker-compose exec postgres psql -U postgres -d gospace -c "SELECT 1;"
```

**Problem**: Cannot access application

```bash
# Check if application is listening
docker-compose exec app netstat -tlnp | grep 8080

# Check firewall rules
sudo ufw status  # Linux
```

## Health Checks

### Docker

```bash
# Check container health
docker-compose ps

# Test application endpoint
curl http://localhost:8080

# Test database connection
docker-compose exec postgres pg_isready -U postgres
```

### Kubernetes

```bash
# Check pod health
kubectl get pods -n gospace

# Check readiness/liveness probes
kubectl describe pod <pod-name> -n gospace | grep -A 10 "Liveness\|Readiness"

# Test application endpoint
kubectl port-forward -n gospace svc/gospace-service 8080:8080
curl http://localhost:8080
```

## Monitoring

### Docker

```bash
# View resource usage
docker stats

# View logs in real-time
docker-compose logs -f --tail=100
```

### Kubernetes

```bash
# View resource usage
kubectl top pods -n gospace
kubectl top nodes

# View events
kubectl get events -n gospace --sort-by='.lastTimestamp'
```

## Backup and Restore

### PostgreSQL Backup

```bash
# Docker
docker-compose exec postgres pg_dump -U postgres gospace > backup.sql

# Kubernetes
kubectl exec -n gospace <postgres-pod> -- pg_dump -U postgres gospace > backup.sql
```

### PostgreSQL Restore

```bash
# Docker
docker-compose exec -T postgres psql -U postgres gospace < backup.sql

# Kubernetes
kubectl exec -i -n gospace <postgres-pod> -- psql -U postgres gospace < backup.sql
```

## Security Considerations

1. **Change default passwords** in production
2. **Use secrets management** for sensitive data (Kubernetes Secrets, Docker Secrets)
3. **Enable SSL/TLS** for database connections in production
4. **Use network policies** in Kubernetes to restrict traffic
5. **Regularly update** base images and dependencies
6. **Scan images** for vulnerabilities using tools like Trivy

## Performance Tuning

### Docker

- Adjust resource limits in docker-compose.yml
- Use volume mounts for better I/O performance
- Enable BuildKit for faster builds

### Kubernetes

- Adjust resource requests/limits in k8s-deployment.yaml
- Use Horizontal Pod Autoscaler for auto-scaling
- Configure PersistentVolume with appropriate storage class
- Use node affinity for optimal pod placement

## Support

For issues and questions:
- Check logs first
- Review this troubleshooting guide
- Check GitHub issues
- Contact support team