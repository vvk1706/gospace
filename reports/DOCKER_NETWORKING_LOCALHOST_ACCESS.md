# Docker Container Localhost Access Troubleshooting
**Date:** 2026-05-13  
**Issue:** Connecting Docker containers (Portainer, pgAdmin) to Kubernetes services via localhost

---

## Problem Statement

After recreating the Kind Kubernetes cluster, both Portainer and pgAdmin containers need to connect to services running in the cluster:
- **pgAdmin** needs to connect to PostgreSQL on port 30432
- **Portainer** needs to connect to Portainer Agent on port 30778

### Initial Attempts

**User tried:** Using `localhost` or `127.0.0.1` in connection strings
**Result:** Connection refused

### Root Cause

Docker containers running in bridge/custom network mode cannot access the host's `localhost`. When a container tries to connect to `127.0.0.1`, it connects to itself, not the host machine.

---

## Investigation Process

### 1. Network Mode Check

```bash
docker inspect portainer --format '{{.HostConfig.NetworkMode}}'
# Output: bridge

docker inspect pgadmin --format '{{.HostConfig.NetworkMode}}'
# Output: gospace_default
```

**Finding:** Both containers are in bridge mode, not host mode.

### 2. Attempted Solution: Connect to Host Network

```bash
docker network connect host portainer
docker network connect host pgadmin
```

**Result:** ❌ Failed
```
Error: cannot connect container to host network - 
container must be created in host network mode
```

**Reason:** Host network mode must be set at container creation time, cannot be added later.

### 3. Current Working Solution

Both containers are connected to the `mybridge` network along with the Kind control-plane:

```bash
docker network inspect mybridge --format '{{json .Containers}}' | jq
```

**Network Topology:**
- Portainer: `172.20.0.2`
- mk-control-plane: `172.20.0.3`
- pgadmin: `172.20.0.7`
- gospace_db: `172.20.0.8`

**Working Connection:**
- pgAdmin → PostgreSQL: `172.20.0.3:30432` ✅
- Portainer → Agent: `172.20.0.3:30778` ✅

### 4. Explored: host.docker.internal

Attempted to use `host.docker.internal` as a hostname that resolves to the host.

**Finding:** `host.docker.internal` is **not available on Linux Docker** by default. It's a feature of Docker Desktop for macOS and Windows.

### 5. Considered: Manual /etc/hosts Entry

**User Concern:** ❌ "What if host IP changes?"
**Valid Issue:** 
- DHCP can change host IP
- Network changes affect IP
- Requires manual updates
- Not a reliable solution

### 6. Concern: Kind Control-Plane IP Stability

**User Question:** "What if I restart Kind control plane and it changes IP address?"

**Current IP:** `172.20.0.3/16` (dynamically assigned)

**Issue:** IP might change on container restart/reconnect.

---

## Solution Options Summary

### Option 1: Use Current Dynamic IP
- pgAdmin Host: `172.20.0.3`
- Portainer URL: `172.20.0.3:30778`
- **Pros:** Works immediately
- **Cons:** IP may change

### Option 2: Assign Static IP (Recommended)
```bash
docker network disconnect mybridge mk-control-plane
docker network connect mybridge mk-control-plane --ip 172.20.0.10
```
- **Pros:** Stable IP, survives restarts
- **Cons:** Brief downtime during reconnection

### Option 3: Docker Compose with extra_hosts
- Use DNS-like hostname
- **Pros:** More maintainable
- **Cons:** Requires container recreation

### Option 4: Host Network Mode
- **Pros:** Direct localhost access
- **Cons:** Loses configurations, port conflicts

### Option 5: Deploy in Kubernetes
- **Pros:** Native networking
- **Cons:** Different architecture

---

## Recommended Solution

**Assign static IP to Kind control-plane:**

```bash
docker network disconnect mybridge mk-control-plane
docker network connect mybridge mk-control-plane --ip 172.20.0.10
```

**Then configure:**
- pgAdmin Host: `172.20.0.10:30432`
- Portainer URL: `172.20.0.10:30778`

---

## Key Learnings

1. `host.docker.internal` not available on Linux Docker
2. Containers in bridge mode cannot access host localhost
3. Dynamic IPs change on restart - use static IPs
4. Static IP provides best balance of reliability and ease

---

**Status:** Pending implementation of static IP solution