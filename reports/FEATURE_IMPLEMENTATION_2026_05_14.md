# Agent/Tool Source Code Storage & Edit Functionality Implementation
**Date:** 2026-05-14  
**Feature:** Add source code storage and edit functionality for AI agents and tools  
**Status:** ✅ Completed and Deployed

---

## Overview

Implemented comprehensive functionality to store, view, and edit source code for AI agents and tools directly in the application database. This enables users to manage their agent/tool implementations alongside metadata.

---

## Requirements

User requested:
1. Store agent/tool source code in database
2. Add functionality to edit agent and tool records
3. Save changes back to database
4. Deploy to both Docker and Kubernetes environments

---

## Implementation Details

### 1. Database Schema Changes

#### Agent Model (`models/agent.go`)
```go
type Agent struct {
    ID          uint           `gorm:"primarykey" json:"id"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
    Name        string         `gorm:"not null;size:255" json:"name" form:"name" binding:"required"`
    Description string         `gorm:"type:text" json:"description" form:"description" binding:"required"`
    Category    string         `gorm:"size:100" json:"category" form:"category" binding:"required"`
    Version     string         `gorm:"size:50" json:"version" form:"version"`
    Author      string         `gorm:"size:255" json:"author" form:"author"`
    Repository  string         `gorm:"size:500" json:"repository" form:"repository"`
    Tags        string         `gorm:"type:text" json:"tags" form:"tags"`
    Status      string         `gorm:"size:50;default:'active'" json:"status" form:"status"`
    SourceCode  string         `gorm:"type:text" json:"source_code" form:"source_code"` // NEW
}
```

#### Tool Model (`models/tool.go`)
```go
type Tool struct {
    ID          uint           `gorm:"primarykey" json:"id"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
    Name        string         `gorm:"not null;size:255" json:"name" form:"name" binding:"required"`
    Description string         `gorm:"type:text" json:"description" form:"description" binding:"required"`
    Category    string         `gorm:"size:100" json:"category" form:"category" binding:"required"`
    Version     string         `gorm:"size:50" json:"version" form:"version"`
    Author      string         `gorm:"size:255" json:"author" form:"author"`
    Repository  string         `gorm:"size:500" json:"repository" form:"repository"`
    Tags        string         `gorm:"type:text" json:"tags" form:"tags"`
    Language    string         `gorm:"size:50" json:"language" form:"language"`
    Status      string         `gorm:"size:50;default:'active'" json:"status" form:"status"`
    SourceCode  string         `gorm:"type:text" json:"source_code" form:"source_code"` // NEW
}
```

**Migration:** AutoMigrate automatically adds the new `source_code` column on application startup.

---

### 2. Handler Functions

#### Agent Handlers (`handlers/agents.go`)

**EditAgentForm** - Display edit form
```go
func (h *Handler) EditAgentForm(c *gin.Context) {
    id := c.Param("id")
    agentID, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.HTML(http.StatusBadRequest, "agents.html", gin.H{
            "title": "AI Agents Repository",
            "error": "Invalid agent ID",
        })
        return
    }

    var agent models.Agent
    if err := h.DB.First(&agent, agentID).Error; err != nil {
        c.HTML(http.StatusNotFound, "agents.html", gin.H{
            "title": "AI Agents Repository",
            "error": "Agent not found",
        })
        return
    }

    c.HTML(http.StatusOK, "edit_agent.html", gin.H{
        "title": "Edit AI Agent",
        "agent": agent,
    })
}
```

**UpdateAgent** - Save changes
```go
func (h *Handler) UpdateAgent(c *gin.Context) {
    id := c.Param("id")
    agentID, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.HTML(http.StatusBadRequest, "edit_agent.html", gin.H{
            "title": "Edit AI Agent",
            "error": "Invalid agent ID",
        })
        return
    }

    var agent models.Agent
    if err := h.DB.First(&agent, agentID).Error; err != nil {
        c.HTML(http.StatusNotFound, "edit_agent.html", gin.H{
            "title": "Edit AI Agent",
            "error": "Agent not found",
        })
        return
    }

    if err := c.ShouldBind(&agent); err != nil {
        c.HTML(http.StatusBadRequest, "edit_agent.html", gin.H{
            "title": "Edit AI Agent",
            "error": "Invalid form data: " + err.Error(),
            "agent": agent,
        })
        return
    }

    if err := h.DB.Save(&agent).Error; err != nil {
        c.HTML(http.StatusInternalServerError, "edit_agent.html", gin.H{
            "title": "Edit AI Agent",
            "error": "Failed to update agent: " + err.Error(),
            "agent": agent,
        })
        return
    }

    c.Redirect(http.StatusSeeOther, "/agents")
}
```

#### Tool Handlers (`handlers/tools.go`)

Similar implementations for `EditToolForm` and `UpdateTool`.

---

### 3. Templates

#### New Templates Created

**templates/edit_agent.html**
- Full edit form with all agent fields
- Monospace textarea for source code (15 rows)
- Pre-populated with existing values
- Update and Cancel buttons

**templates/edit_tool.html**
- Full edit form with all tool fields
- Monospace textarea for source code (15 rows)
- Pre-populated with existing values
- Update and Cancel buttons

#### Updated Templates

**templates/add_agent.html**
- Added source code textarea field

**templates/add_tool.html**
- Added source code textarea field

**templates/agents.html**
- Added edit button (✏️) in actions column
- Links to `/agents/:id/edit`

**templates/tools.html**
- Added edit button (✏️) in actions column
- Links to `/tools/:id/edit`

---

### 4. Routes

#### New Routes Added (`main.go`)

```go
// AI Agents routes
router.GET("/agents", h.ListAgents)
router.GET("/agents/add", h.AddAgentForm)
router.POST("/agents/add", h.CreateAgent)
router.GET("/agents/:id/edit", h.EditAgentForm)      // NEW
router.POST("/agents/:id/edit", h.UpdateAgent)       // NEW
router.POST("/agents/:id/delete", h.DeleteAgent)

// AI Tools routes
router.GET("/tools", h.ListTools)
router.GET("/tools/add", h.AddToolForm)
router.POST("/tools/add", h.CreateTool)
router.GET("/tools/:id/edit", h.EditToolForm)        // NEW
router.POST("/tools/:id/edit", h.UpdateTool)         // NEW
router.POST("/tools/:id/delete", h.DeleteTool)
```

---

## Deployment Process

### 1. Docker Deployment

```bash
# Stop existing containers
docker compose down

# Rebuild and start with new code
docker compose up -d --build

# Connect to mybridge network
docker network connect mybridge gospace_db

# Verify
docker logs gospace --tail 20
```

**Result:** ✅ Running on port 8080 with new routes active

---

### 2. Kubernetes Deployment

```bash
# Build new image without cache
docker build --no-cache -t localhost:5000/gospace:v2 .

# Tag for Docker Hub
docker tag localhost:5000/gospace:v2 vvk17/gospace:latest

# Push to Docker Hub
docker push vvk17/gospace:latest

# Rollout restart DaemonSet
kubectl rollout restart daemonset gospace -n gospace-app

# Wait for pods to be ready
kubectl get pods -n gospace-app -w

# Verify new templates in pods
kubectl exec -n gospace-app gospace-gcwqh -- ls -la /root/templates/
# Output shows: edit_agent.html, edit_tool.html

# Verify new routes registered
kubectl logs gospace-gcwqh -n gospace-app | grep edit
# Output shows: GET /agents/:id/edit, POST /agents/:id/edit, etc.
```

**Result:** ✅ All 4 pods running with new image and functionality

---

## Verification

### Template Verification
```bash
kubectl exec -n gospace-app gospace-gcwqh -- ls -la /root/templates/
```
**Output:**
```
-rw-rw-r--    1 root     root          5352 May 14 09:03 edit_agent.html
-rw-rw-r--    1 root     root          6427 May 14 09:04 edit_tool.html
-rw-rw-r--    1 root     root          4223 May 14 09:03 add_agent.html
-rw-rw-r--    1 root     root          4937 May 14 09:03 add_tool.html
```

### Route Verification
```bash
kubectl logs gospace-gcwqh -n gospace-app | grep "agents\|tools"
```
**Output:**
```
[GIN-debug] GET    /agents                   --> ...
[GIN-debug] GET    /agents/add               --> ...
[GIN-debug] POST   /agents/add               --> ...
[GIN-debug] GET    /agents/:id/edit          --> ...
[GIN-debug] POST   /agents/:id/edit          --> ...
[GIN-debug] POST   /agents/:id/delete        --> ...
[GIN-debug] GET    /tools                    --> ...
[GIN-debug] GET    /tools/add                --> ...
[GIN-debug] POST   /tools/add                --> ...
[GIN-debug] GET    /tools/:id/edit           --> ...
[GIN-debug] POST   /tools/:id/edit           --> ...
[GIN-debug] POST   /tools/:id/delete         --> ...
```

---

## Git Repository

### Commit Details
```bash
git add -A
git commit -m "feat: Add source code storage and edit functionality for agents and tools

- Add SourceCode field to Agent and Tool models for storing code in database
- Implement edit handlers (EditAgentForm, UpdateAgent, EditToolForm, UpdateTool)
- Create edit templates (edit_agent.html, edit_tool.html) with monospace code textarea
- Update add templates to include source code field
- Add edit buttons (✏️) to agents and tools list pages
- Register new routes for editing agents and tools
- Document cluster connectivity fix with static IP assignment (172.20.0.10)

This enables users to store, view, and edit agent/tool source code directly in the application."

git push origin main
```

**Commit:** c99aa35  
**Files Changed:** 12 files  
**Insertions:** 564 lines  
**Repository:** https://github.com/vvk1706/gospace

---

## Access Information

### Docker Environment
- **URL:** http://localhost:8080
- **Database:** PostgreSQL on gospace_db container
- **Network:** gospace_default + mybridge

### Kubernetes Environment
- **URL:** http://localhost:30081
- **Alternative:** http://172.20.0.10:30081
- **Namespace:** gospace-app
- **Pods:** 4 (DaemonSet across all worker nodes)
- **Database:** PostgreSQL in gospace-db namespace

---

## User Guide

### How to Edit an Agent or Tool

1. **Navigate to List Page**
   - Go to http://localhost:30081/agents or /tools

2. **Click Edit Button**
   - Click the ✏️ icon next to the agent/tool you want to edit

3. **Edit Fields**
   - Modify any field including the source code textarea
   - Source code textarea uses monospace font for better readability

4. **Save Changes**
   - Click "Update Agent" or "Update Tool" button
   - Changes are saved to PostgreSQL database
   - Redirects back to list page

5. **Verify**
   - Source code persists across application restarts
   - Can be edited again at any time

---

## Technical Notes

### Database Schema Migration
- GORM AutoMigrate automatically adds the `source_code` column
- Existing records will have NULL/empty source code initially
- No manual migration required

### Source Code Storage
- Field type: TEXT (unlimited length in PostgreSQL)
- Supports large code files
- UTF-8 encoding for special characters

### Form Binding
- Uses Gin's `ShouldBind` for automatic form parsing
- Validates required fields
- Preserves existing data on validation errors

---

## Related Documentation

- [Cluster Restart Fix](./CLUSTER_RESTART_FIX_2026_05_14.md) - Static IP assignment for mk-control-plane
- [Docker Networking](./DOCKER_NETWORKING_LOCALHOST_ACCESS.md) - Container networking troubleshooting
- [Kind Troubleshooting](./TROUBLESHOOTING_KIND_2026_05_13.md) - Cluster setup and issues

---

## Status Summary

✅ **Database Schema:** Updated with SourceCode field  
✅ **Handlers:** Edit and Update functions implemented  
✅ **Templates:** Edit forms created, list pages updated  
✅ **Routes:** New edit routes registered  
✅ **Docker Deployment:** Rebuilt and running  
✅ **Kubernetes Deployment:** All pods updated with new image  
✅ **Git Repository:** Changes committed and pushed  
✅ **Documentation:** Complete implementation guide created

---

**Implementation Date:** 2026-05-14  
**Implemented By:** Bob (AI Assistant)  
**Tested On:** Docker (localhost:8080) and Kubernetes (localhost:30081)  
**Status:** Production Ready ✅