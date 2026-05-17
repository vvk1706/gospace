# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

## [1.1.0] - 2026-05-17

### Added
- AI Agents repository management feature
- AI Tools repository management feature
- Horizontal layout for home page feature cards
- 2-column horizontal form layout for agent/tool add/edit forms
- Comprehensive UI modernization report (reports/UI_MODERNIZATION_2026_05_17.md)
- Consistent 6-item navigation menu across all pages
- Footer added to calculator_history.html

### Changed
- Home page feature cards now display horizontally using flexbox instead of vertical grid
- Form container max-width increased from 600px to 1800px (3x wider)
- All form fields (input, select, textarea) now have consistent styling
- Navigation menu standardized across all pages with 6 items (Home, Calculator, Contact, View Contacts, AI Agents, AI Tools)
- Logo name unified to "GoSpace App" across all templates
- Footer text unified to "GoSpace App" across all templates
- Calculator history page styling updated to match other pages
- Docker Compose configuration updated to use pre-built image instead of building
- Kubernetes deployment updated to use standardized image name
- Agent and tool forms now display fields in 2-column layout for better space utilization

### Fixed
- Inconsistent navigation menu items across different pages
- Missing AI Agents and AI Tools links on some pages
- Missing footer on calculator_history.html
- Inconsistent logo naming (GoSpace vs GoSpace App)

### Deployment
- Standardized Docker image name to `gospace:latest`
- Removed old images: `gospace-app:latest`, `gospace-k:latest`, `localhost:5000/gospace:v2`
- Updated docker-compose.yml to use `gospace:latest` image
- Updated k8s-deployment-new.yaml to use `vvk17/gospace:latest` image
- Pushed images to both local registry (localhost:5000/gospace:latest) and Docker Hub (vvk17/gospace:latest)

### Technical Details
- CSS: Added `.form-horizontal` class with grid layout for 2-column forms
- CSS: Added `.form-group-full` class for full-width form fields
- CSS: Updated `.features` from grid to flexbox layout
- CSS: Updated `.feature-card` with flex sizing
- CSS: Increased `.form-container` max-width to 1800px
- Templates: Updated 11 template files for consistency and new features

## [1.0.0] - Previous Release

### Added
- Comprehensive test coverage for calculator redirect behavior
- Tests for delete-only workflow in calculator history

### Changed
- Calculator form now redirects to history page after successful calculation instead of displaying result inline
- Simplified calculator history interface - removed inline editing functionality
- Calculator history now supports delete-only operations
- API documentation updated to reflect actual behavior

### Removed
- Client-side JavaScript files (`calculator.js`, `contact.js`, `main.js`)
- Inline editing functionality from calculator history
- `EditCalculatorHistory` handler and route
- Tests for edit functionality (229 lines removed)
- CSRF protection and session management (removed for simplicity)
- `SESSION_SECRET` and `CSRF_SECRET` configuration

### Breaking Changes
- POST `/calculator/history/:id/edit` endpoint removed
- Calculator POST endpoint now returns `303 See Other` redirect instead of HTML with result

## Migration Guide

### For Users
- No action required - the application will work with the simplified interface
- Edit functionality has been removed; use delete and re-create instead

### For Developers
- Remove `SESSION_SECRET` and `CSRF_SECRET` from `.env` file if present
- Remove any CSRF token fields from custom forms
- Tests no longer need session middleware or CSRF configuration
- Remove any code that relied on the edit endpoint

### Rationale
The removal of JavaScript, edit functionality, and CSRF simplifies the application architecture:
- Reduces client-side complexity
- Simplifies server-side middleware stack
- Eliminates potential XSS vulnerabilities from client-side code
- Simplifies maintenance and testing
- Follows server-side rendering best practices
- Reduces dependencies and attack surface