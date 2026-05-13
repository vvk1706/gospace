# Kind Kubernetes Cluster - "mk"

This directory contains configuration and scripts for deploying a Kind (Kubernetes in Docker) cluster named "mk" with 1 control plane node and 4 worker nodes.

## Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Cluster Configuration](#cluster-configuration)
- [Deployment](#deployment)
- [Management](#management)
- [Troubleshooting](#troubleshooting)
- [Advanced Usage](#advanced-usage)

## Overview

**Kind** (Kubernetes IN Docker) is a tool for running local Kubernetes clusters using Docker container "nodes". It's designed for testing Kubernetes itself, but can also be used for local development or CI.

### Cluster Specifications

- **Cluster Name**: mk
- **Control Plane Nodes**: 1
- **Worker Nodes**: 4
- **Kubernetes Version**: Latest (determined by Kind version)
- **Container Runtime**: Docker
- **Network Plugin**: Default CNI (Kindnet)

### Port Mappings

The control plane node exposes the following ports:

- `80` → Container port 80 (HTTP)
- `443` → Container port 443 (HTTPS)
- `30000` → Container port 30000 (Custom NodePort)
- `30778` → Container port 30778 (Portainer Agent)

## Prerequisites

### Required Software

1. **Docker** (version 20.10 or later)
   ```bash
   docker --version
   ```

2. **Kind** (version 0.20.0 or later)
   ```bash
   kind version
   ```

3. **kubectl** (matching your Kubernetes version)
   ```bash
   kubectl version --client
   ```

### Installing Kind

If Kind is not installed, you can install it using:

```bash
# Download Kind binary
curl -Lo ~/bin/kind https://kind.sigs.k8s.io/dl/v0.20.0/kind-linux-amd64

# Make it executable
chmod +x ~/bin/kind

# Add to PATH (add this to your ~/.bashrc or ~/.zshrc)
export PATH=$PATH:~/bin

# Verify installation
kind version
```

### Installing kubectl

```bash
# Download kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"

# Make it executable
chmod +x kubectl

# Move to PATH
mkdir -p ~/bin
mv kubectl ~/bin/

# Verify installation
kubectl version --client
```

## Quick Start

### 1. Deploy the Cluster

```bash
cd infrastructure/kubernetes/clusters/Kind
chmod +x deploy.sh
./deploy.sh
```

### 2. Verify the Cluster

```bash
# Check cluster status
kubectl cluster-info

# List all nodes
kubectl get nodes

# Check node details
kubectl get nodes -o wide
```

### 3. Deploy Applications

```bash
# Example: Deploy a test application
kubectl create deployment nginx --image=nginx
kubectl expose deployment nginx --port=80 --type=NodePort
```

## Cluster Configuration

### Configuration File: `kind-config.yaml`

The cluster is configured using the `kind-config.yaml` file with the following key settings:

#### Nodes Configuration

```yaml
nodes:
  - role: control-plane    # 1 control plane node
  - role: worker          # Worker node 1
  - role: worker          # Worker node 2
  - role: worker          # Worker node 3
  - role: worker          # Worker node 4
```

#### Networking

- **Pod Subnet**: `10.244.0.0/16`
- **Service Subnet**: `10.96.0.0/12`
- **API Server**: `127.0.0.1:6443`

#### Node Labels

Worker nodes are labeled for easy identification:
- `node-type: worker`
- `worker-id: "1"` through `worker-id: "4"`

#### Feature Gates

- `EphemeralContainers: true` - Enables ephemeral containers for debugging

## Deployment

### Using the Deployment Script

The `deploy.sh` script automates the cluster creation process:

```bash
./deploy.sh
```

**Features:**
- Checks for required dependencies
- Detects existing clusters and prompts for recreation
- Creates cluster with specified configuration
- Displays cluster information after creation
- Provides helpful next steps

### Manual Deployment

If you prefer to deploy manually:

```bash
# Create cluster
kind create cluster --config kind-config.yaml

# Verify cluster
kubectl cluster-info
kubectl get nodes
```

### Post-Deployment Steps

1. **Verify all nodes are Ready**:
   ```bash
   kubectl get nodes
   ```

2. **Check system pods**:
   ```bash
   kubectl get pods -n kube-system
   ```

3. **Set up kubectl context** (if needed):
   ```bash
   kubectl config use-context kind-mk
   ```

## Management

### Viewing Cluster Information

```bash
# Get cluster info
kubectl cluster-info

# List all clusters
kind get clusters

# Get nodes
kubectl get nodes -o wide

# Get all resources
kubectl get all --all-namespaces
```

### Accessing the Cluster

```bash
# Get kubeconfig
kind get kubeconfig --name mk

# Export kubeconfig
kind export kubeconfig --name mk
```

### Deleting the Cluster

#### Using the Deletion Script

```bash
./delete.sh
```

#### Manual Deletion

```bash
kind delete cluster --name mk
```

### Stopping and Starting

**Note**: Kind clusters cannot be stopped and started like traditional VMs. You must delete and recreate them.

## Troubleshooting

### Common Issues

#### 1. Cluster Creation Fails

**Problem**: Cluster creation fails with network errors

**Solution**:
```bash
# Check Docker is running
docker ps

# Check Docker network
docker network ls

# Restart Docker if needed
sudo systemctl restart docker
```

#### 2. Nodes Not Ready

**Problem**: Nodes show "NotReady" status

**Solution**:
```bash
# Check node status
kubectl describe nodes

# Check system pods
kubectl get pods -n kube-system

# Wait for all pods to be running
kubectl wait --for=condition=Ready pods --all -n kube-system --timeout=300s
```

#### 3. Cannot Connect to Cluster

**Problem**: kubectl cannot connect to the cluster

**Solution**:
```bash
# Verify cluster exists
kind get clusters

# Update kubeconfig
kind export kubeconfig --name mk

# Check current context
kubectl config current-context

# Switch to correct context
kubectl config use-context kind-mk
```

#### 4. Port Conflicts

**Problem**: Ports 80, 443, or others are already in use

**Solution**:
```bash
# Check what's using the ports
sudo lsof -i :80
sudo lsof -i :443

# Stop conflicting services or modify kind-config.yaml port mappings
```

### Debugging Commands

```bash
# View cluster logs
kind export logs --name mk

# Describe a node
kubectl describe node <node-name>

# Get events
kubectl get events --all-namespaces --sort-by='.lastTimestamp'

# Check Docker containers
docker ps -a | grep mk
```

## Advanced Usage

### Installing Ingress Controller

```bash
# Install NGINX Ingress Controller
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

# Wait for ingress controller to be ready
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s
```

### Installing Portainer Agent

```bash
# Apply Portainer agent configuration
kubectl apply -f ../../../portainer-agent.yaml

# Get agent service URL
kubectl get svc -n portainer portainer-agent

# Access via NodePort 30778
```

### Load Balancer with MetalLB

```bash
# Install MetalLB
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.7/config/manifests/metallb-native.yaml

# Configure IP address pool (adjust as needed)
kubectl apply -f - <<EOF
apiVersion: metallb.io/v1beta1
kind: IPAddressPool
metadata:
  name: example
  namespace: metallb-system
spec:
  addresses:
  - 172.18.255.200-172.18.255.250
---
apiVersion: metallb.io/v1beta1
kind: L2Advertisement
metadata:
  name: empty
  namespace: metallb-system
EOF
```

### Persistent Storage

Kind supports local path provisioner by default:

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
  storageClassName: standard
```

### Multi-Cluster Setup

To run multiple Kind clusters:

```bash
# Create additional cluster with different name
kind create cluster --name mk-dev --config kind-config-dev.yaml

# Switch between clusters
kubectl config use-context kind-mk
kubectl config use-context kind-mk-dev
```

### Custom Docker Images

Load local Docker images into the cluster:

```bash
# Build your image
docker build -t my-app:latest .

# Load into Kind cluster
kind load docker-image my-app:latest --name mk

# Use in deployments
kubectl create deployment my-app --image=my-app:latest
```

### Cluster Backup and Restore

```bash
# Export cluster configuration
kubectl get all --all-namespaces -o yaml > cluster-backup.yaml

# Restore to new cluster
kubectl apply -f cluster-backup.yaml
```

## Best Practices

1. **Resource Limits**: Always set resource requests and limits for production-like testing
2. **Namespaces**: Use namespaces to organize resources
3. **Labels**: Use consistent labeling for easy resource management
4. **Monitoring**: Install monitoring tools (Prometheus, Grafana) for observability
5. **Cleanup**: Delete the cluster when not in use to free resources

## Additional Resources

- [Kind Documentation](https://kind.sigs.k8s.io/)
- [Kubernetes Documentation](https://kubernetes.io/docs/)
- [kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)
- [Kind GitHub Repository](https://github.com/kubernetes-sigs/kind)

## Support

For issues specific to this cluster setup, please check:
1. This README's troubleshooting section
2. Kind's official documentation
3. Project's issue tracker

---

**Last Updated**: 2026-05-13  
**Maintained By**: DevOps Team