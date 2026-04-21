# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

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