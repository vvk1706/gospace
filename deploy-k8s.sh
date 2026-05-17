#!/bin/bash

# GoSpace Kubernetes Deployment Script
# This script deploys the GoSpace application to Kubernetes

set -e

echo "🚀 GoSpace Kubernetes Deployment"
echo "=================================="
echo ""

# Check if kubectl is available
if ! command -v kubectl &> /dev/null; then
    echo "❌ kubectl not found. Please install kubectl first."
    exit 1
fi

# Build and push Docker image
echo "📦 Building Docker image..."
docker build -t gospace:latest -t localhost:5000/gospace:latest -t vvk17/gospace:latest .

echo "📤 Pushing to registries..."
docker push localhost:5000/gospace:latest
docker push vvk17/gospace:latest
echo "✅ Docker images built and pushed"

echo ""
echo "🔄 Deploying to Kubernetes..."

# Create namespace and deploy PostgreSQL
echo "   - Creating namespace and PostgreSQL..."
kubectl apply -f k8s-postgres-new.yaml

# Wait for PostgreSQL to be ready
echo "   - Waiting for PostgreSQL to be ready..."
kubectl wait --for=condition=ready pod -l app=postgres -n gospace-db --timeout=120s

# Deploy application
echo "   - Deploying GoSpace application..."
kubectl apply -f k8s-deployment-new.yaml

# Wait for application to be ready
echo "   - Waiting for application to be ready..."
sleep 10
kubectl get pods -n gospace-app

echo ""
echo "✅ Deployment complete!"
echo ""
echo "📊 Deployment Status:"
kubectl get all -n gospace-app

echo ""
echo "🌐 Access Information:"
echo "   - NodePort: http://localhost:30081"
echo "   - Ingress: http://gospace.local (if ingress controller is configured)"
echo ""
echo "📝 Useful commands:"
echo "   - View pods: kubectl get pods -n gospace-app"
echo "   - View logs: kubectl logs -f daemonset/gospace -n gospace-app"
echo "   - Port forward: kubectl port-forward -n gospace-app svc/gospace-service 8080:8080"
echo "   - Delete deployment: kubectl delete namespace gospace-app"
echo ""

# Made with Bob
