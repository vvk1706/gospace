# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added
- CSRF protection for all POST endpoints using `gin-csrf` middleware
- Session management using `gin-contrib/sessions`
- Security configuration via environment variables (`SESSION_SECRET`, `CSRF_SECRET`)
- Comprehensive test coverage for calculator redirect behavior
- Tests for delete-only workflow in calculator history

### Changed
- Calculator form now redirects to history page after successful calculation instead of displaying result inline
- Simplified calculator history interface - removed inline editing functionality
- Calculator history now supports delete-only operations
- All forms now include CSRF tokens for security
- API documentation updated to reflect actual behavior and security requirements

### Removed
- Client-side JavaScript files (`calculator.js`, `contact.js`, `main.js`)
- Inline editing functionality from calculator history
- `EditCalculatorHistory` handler and route
- Tests for edit functionality (229 lines removed)

### Security
- Added CSRF protection to prevent cross-site request forgery attacks
- All POST requests now require valid CSRF tokens
- Session secrets should be changed from defaults in production

### Breaking Changes
- POST `/calculator/history/:id/edit` endpoint removed
- Calculator POST endpoint now returns `303 See Other` redirect instead of HTML with result
- All POST forms require `_csrf` hidden field with valid token

## Migration Guide

### For Users
- No action required - the application will work with the simplified interface
- Edit functionality has been removed; use delete and re-create instead

### For Developers
- Update `.env` file with `SESSION_SECRET` and `CSRF_SECRET` values
- All POST forms must include CSRF token: `<input type="hidden" name="_csrf" value="{{ .csrf }}">`
- Tests must include session middleware or disable CSRF validation
- Remove any code that relied on the edit endpoint

### Rationale
The removal of JavaScript and edit functionality simplifies the application architecture:
- Reduces client-side complexity
- Improves security posture with server-side validation only
- Eliminates potential XSS vulnerabilities from client-side code
- Simplifies maintenance and testing
- Follows server-side rendering best practices