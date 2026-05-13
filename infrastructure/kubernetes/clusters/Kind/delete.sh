#!/bin/bash

# Kind Cluster Deletion Script
# This script deletes the Kind Kubernetes cluster

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration
CLUSTER_NAME="mk"

echo -e "${YELLOW}========================================${NC}"
echo -e "${YELLOW}Kind Cluster Deletion Script${NC}"
echo -e "${YELLOW}========================================${NC}"
echo ""

# Check if Kind is installed
if ! command -v kind &> /dev/null; then
    echo -e "${RED}Error: Kind is not installed${NC}"
    exit 1
fi

# Check if cluster exists
if ! kind get clusters 2>/dev/null | grep -q "^${CLUSTER_NAME}$"; then
    echo -e "${YELLOW}Cluster '${CLUSTER_NAME}' does not exist${NC}"
    exit 0
fi

# Confirm deletion
echo -e "${YELLOW}This will delete the cluster '${CLUSTER_NAME}' and all its resources${NC}"
read -p "Are you sure you want to continue? (y/N): " -n 1 -r
echo

if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo -e "${GREEN}Deletion cancelled${NC}"
    exit 0
fi

# Delete the cluster
echo -e "${YELLOW}Deleting cluster '${CLUSTER_NAME}'...${NC}"
kind delete cluster --name "${CLUSTER_NAME}"

if [ $? -eq 0 ]; then
    echo ""
    echo -e "${GREEN}✓ Cluster '${CLUSTER_NAME}' deleted successfully${NC}"
else
    echo -e "${RED}Failed to delete cluster${NC}"
    exit 1
fi

# Made with Bob
