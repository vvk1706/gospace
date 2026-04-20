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

# Check if Docker image exists
if ! docker image inspect gospace:latest &> /dev/null; then
    echo "📦 Building Docker image..."
    docker build -t gospace:latest .
else
    echo "✅ Docker image gospace:latest found"
fi

echo ""
echo "🔄 Deploying to Kubernetes..."

# Create namespace and deploy PostgreSQL
echo "   - Creating namespace and PostgreSQL..."
kubectl apply -f k8s-postgres.yaml

# Wait for PostgreSQL to be ready
echo "   - Waiting for PostgreSQL to be ready..."
kubectl wait --for=condition=ready pod -l app=postgres -n gospace --timeout=120s

# Deploy application
echo "   - Deploying GoSpace application..."
kubectl apply -f k8s-deployment.yaml

# Wait for application to be ready
echo "   - Waiting for application to be ready..."
kubectl wait --for=condition=ready pod -l app=gospace -n gospace --timeout=120s

echo ""
echo "✅ Deployment complete!"
echo ""
echo "📊 Deployment Status:"
kubectl get all -n gospace

echo ""
echo "🌐 Access Information:"
echo "   - NodePort: http://localhost:30080"
echo "   - Ingress: http://gospace.local (if ingress controller is configured)"
echo ""
echo "📝 Useful commands:"
echo "   - View pods: kubectl get pods -n gospace"
echo "   - View logs: kubectl logs -f deployment/gospace -n gospace"
echo "   - Port forward: kubectl port-forward -n gospace svc/gospace-service 8080:8080"
echo "   - Delete deployment: kubectl delete namespace gospace"
echo ""

# Made with Bob
