#!/bin/bash

# Kind Cluster Deployment Script
# This script creates a Kind Kubernetes cluster with 1 control plane and 4 worker nodes

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration
CLUSTER_NAME="mk"
CONFIG_FILE="kind-config.yaml"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Kind Cluster Deployment Script${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""

# Check if Kind is installed
if ! command -v kind &> /dev/null; then
    echo -e "${RED}Error: Kind is not installed${NC}"
    echo "Please install Kind first:"
    echo "  curl -Lo ~/bin/kind https://kind.sigs.k8s.io/dl/v0.20.0/kind-linux-amd64"
    echo "  chmod +x ~/bin/kind"
    echo "  export PATH=\$PATH:~/bin"
    exit 1
fi

echo -e "${GREEN}✓ Kind version: $(kind version)${NC}"
echo ""

# Check if kubectl is installed
if ! command -v kubectl &> /dev/null; then
    echo -e "${YELLOW}Warning: kubectl is not installed${NC}"
    echo "You may need to install kubectl to interact with the cluster"
fi

# Check if cluster already exists
if kind get clusters 2>/dev/null | grep -q "^${CLUSTER_NAME}$"; then
    echo -e "${YELLOW}Cluster '${CLUSTER_NAME}' already exists${NC}"
    read -p "Do you want to delete and recreate it? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo -e "${YELLOW}Deleting existing cluster...${NC}"
        kind delete cluster --name "${CLUSTER_NAME}"
        echo -e "${GREEN}✓ Cluster deleted${NC}"
    else
        echo -e "${YELLOW}Keeping existing cluster${NC}"
        exit 0
    fi
fi

# Create the cluster
echo -e "${GREEN}Creating Kind cluster '${CLUSTER_NAME}'...${NC}"
echo "Configuration file: ${CONFIG_FILE}"
echo ""

cd "${SCRIPT_DIR}"
kind create cluster --config "${CONFIG_FILE}"

if [ $? -eq 0 ]; then
    echo ""
    echo -e "${GREEN}========================================${NC}"
    echo -e "${GREEN}✓ Cluster created successfully!${NC}"
    echo -e "${GREEN}========================================${NC}"
    echo ""
    
    # Display cluster info
    echo -e "${GREEN}Cluster Information:${NC}"
    echo "  Name: ${CLUSTER_NAME}"
    echo "  Nodes: 1 control-plane + 4 workers"
    echo ""
    
    # Get nodes
    echo -e "${GREEN}Cluster Nodes:${NC}"
    kubectl get nodes -o wide
    echo ""
    
    # Display context
    echo -e "${GREEN}Current Context:${NC}"
    kubectl config current-context
    echo ""
    
    # Display useful commands
    echo -e "${GREEN}Useful Commands:${NC}"
    echo "  View nodes:           kubectl get nodes"
    echo "  View all resources:   kubectl get all --all-namespaces"
    echo "  Delete cluster:       kind delete cluster --name ${CLUSTER_NAME}"
    echo "  Cluster info:         kubectl cluster-info"
    echo ""
    
    echo -e "${GREEN}Next Steps:${NC}"
    echo "  1. Deploy applications to the cluster"
    echo "  2. Install Ingress controller (if needed)"
    echo "  3. Deploy Portainer agent for management"
    echo ""
else
    echo -e "${RED}Failed to create cluster${NC}"
    exit 1
fi

# Made with Bob
