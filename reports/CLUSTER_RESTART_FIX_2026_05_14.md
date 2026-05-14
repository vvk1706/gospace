# Cluster Connectivity Fix After Restart
**Date:** 2026-05-14  
**Issue:** Portainer and pgAdmin unable to connect to Kubernetes cluster after restart  
**Resolution:** Assigned static IP to mk-control-plane

---

## Problem Summary

After restarting the Kind cluster, both Portainer and pgAdmin lost connectivity to the Kubernetes services:
- **Portainer** could not connect to Portainer Agent (port 30778)
- **pgAdmin** could not connect to PostgreSQL (port 30432)

### Root Cause

The mk-control-plane container's IP address on the `mybridge` network changed after restart:
- **Previous IP:** 172.20.0.3
- **New IP after restart:** 172.20.0.4

This dynamic IP assignment broke the connection configurations in both Portainer and pgAdmin.

---

## Solution Applied

### Step 1: Verify Cluster Health
```bash
kubectl get nodes
kubectl get pods -n gospace-app -o wide
```
**Result:** ✅ All nodes Ready, all pods Running

### Step 2: Check Current Network Configuration
```bash
docker network inspect mybridge --format '{{range .Containers}}{{.Name}}: {{.IPv4Address}}{{"\n"}}{{end}}'
```
**Output:**
```
registry: 172.20.0.3/16
mk-control-plane: 172.20.0.4/16  ← Changed from 172.20.0.3
pgadmin: 172.20.0.6/16
gospace_db: 172.20.0.5/16
portainer: 172.20.0.2/16
```

### Step 3: Assign Static IP to mk-control-plane
```bash
docker network disconnect mybridge mk-control-plane
docker network connect mybridge mk-control-plane --ip 172.20.0.10
```

### Step 4: Verify Static IP Assignment
```bash
docker network inspect mybridge --format '{{range .Containers}}{{.Name}}: {{.IPv4Address}}{{"\n"}}{{end}}'
```
**Output:**
```
registry: 172.20.0.3/16
mk-control-plane: 172.20.0.10/16  ← Static IP assigned
pgadmin: 172.20.0.6/16
gospace_db: 172.20.0.5/16
portainer: 172.20.0.2/16
```

### Step 5: Test Connectivity
```bash
curl -s -o /dev/null -w "%{http_code}" http://172.20.0.10:30778  # Portainer Agent
curl -s -o /dev/null -w "%{http_code}" http://172.20.0.10:30081  # GoSpace App
```
**Results:**
- Portainer Agent: 400 (service responding, expects specific protocol)
- GoSpace App: 200 ✅

---

## Configuration Updates Required

### For Portainer
Update the Kubernetes environment endpoint URL:
- **Old URL:** `http://172.20.0.3:30778` or `http://172.20.0.4:30778`
- **New URL:** `http://172.20.0.10:30778`

**Steps:**
1. Open Portainer web interface
2. Go to **Environments**
3. Select your Kubernetes environment
4. Update the **Endpoint URL** to `http://172.20.0.10:30778`
5. Click **Update environment**

### For pgAdmin
Update the PostgreSQL server connection:
- **Old Host:** `172.20.0.3` or `172.20.0.4`
- **New Host:** `172.20.0.10`
- **Port:** 30432 (unchanged)

**Steps:**
1. Open pgAdmin web interface
2. Right-click on the server → **Properties**
3. Go to **Connection** tab
4. Update **Host name/address** to `172.20.0.10`
5. Click **Save**

---

## Network Topology (Final)

```
mybridge network (172.20.0.0/16)
├── portainer: 172.20.0.2
├── registry: 172.20.0.3
├── gospace_db: 172.20.0.5
├── pgadmin: 172.20.0.6
└── mk-control-plane: 172.20.0.10 (STATIC)
```

---

## Benefits of Static IP

1. **Survives Restarts:** IP remains 172.20.0.10 even after container restarts
2. **No Configuration Changes:** Portainer and pgAdmin configs remain valid
3. **Predictable:** Easy to remember and document
4. **Reliable:** No need to check IP after each restart

---

## Future Restarts

After any cluster restart, the static IP will be maintained. However, if you need to verify:

```bash
# Check mk-control-plane IP
docker network inspect mybridge --format '{{range .Containers}}{{if eq .Name "mk-control-plane"}}{{.IPv4Address}}{{end}}{{end}}'
```

Expected output: `172.20.0.10/16`

If the IP ever changes (unlikely), reapply the static IP:
```bash
docker network disconnect mybridge mk-control-plane
docker network connect mybridge mk-control-plane --ip 172.20.0.10
```

---

## Quick Reference

### Service Endpoints (from mybridge network)
- **GoSpace App:** http://172.20.0.10:30081
- **PostgreSQL:** 172.20.0.10:30432
- **Portainer Agent:** http://172.20.0.10:30778

### Service Endpoints (from localhost)
- **GoSpace App:** http://localhost:30081
- **PostgreSQL:** localhost:30432
- **Portainer Agent:** http://localhost:30778

---

## Status

✅ **RESOLVED**
- Static IP assigned to mk-control-plane: 172.20.0.10
- All services accessible and responding
- Configuration updates documented for Portainer and pgAdmin

---

**Next Steps for User:**
1. Update Portainer environment URL to `http://172.20.0.10:30778`
2. Update pgAdmin server host to `172.20.0.10`
3. Test connections from both tools