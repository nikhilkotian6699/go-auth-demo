Sure — I'll update the README so the second developer is **Clooyzi** (the GitHub test/developer account), while keeping **Nikhil** as the primary developer account.

# Go Auth Demo

A simple web-based user authentication project developed as a team collaboration exercise.

The project demonstrates a basic user journey from registration to login, session management, dashboard access, and logout while following a structured GitHub development workflow.

---

## Project Objective

The objective of this project is to build a simple authentication application while practicing:

* Team-based software development
* GitHub collaboration
* Feature-based development
* Pull Requests
* Code reviews
* Automated testing and builds
* Basic CI workflow

---

## Team

| Team Member | Responsibility              |
| ----------- | --------------------------- |
| **Nikhil**  | Go Backend & Authentication |
| **Clooyzi** | HTML Pages & User Sessions  |

---

## Project Scope

The application will support:

```text
Register
   ↓
Login
   ↓
Session
   ↓
Dashboard
   ↓
Logout
```

### Main Features

* User Registration
* User Login
* User Authentication
* User Session
* Protected Dashboard
* Logout
* Automated Testing
* Automated Build

---

## Development Workflow

The project follows a simple team development process:

```text
Task
 ↓
Developer Assignment
 ↓
Feature Development
 ↓
GitHub Pull Request
 ↓
Code Review
 ↓
Automated Test & Build
 ↓
Develop
 ↓
Final Review
 ↓
Main
```

---

## Branches

| Branch      | Purpose                   |
| ----------- | ------------------------- |
| `main`      | Stable and final version  |
| `develop`   | Active development        |
| `feature/*` | Individual developer work |

### Current Feature Branches

```text
feature/auth
feature/session-html
```

---

## Team Responsibilities

### Nikhil

Responsible for:

* Go backend
* User registration
* User login
* Authentication
* Basic testing
* CI setup

### Clooyzi

Responsible for:

* Registration page
* Login page
* Dashboard
* User session
* Logout
* User flow testing

---

## Collaboration Rules

1. Developers work on their assigned feature.
2. Changes are pushed to the respective feature branch.
3. Changes are submitted through a Pull Request.
4. The other developer reviews the work.
5. Automated checks must pass.
6. Approved changes are merged into `develop`.
7. Completed and tested work is eventually merged into `main`.

### Important

**Direct feature development on `main` is not allowed.**

---

## Code Review

Each developer reviews the other's work.

```text
Nikhil's Work → Clooyzi Review

Clooyzi's Work → Nikhil Review
```

Reviews focus on:

* Task completion
* Application functionality
* Code quality
* Testing
* Possible issues
* Requirements

---

## CI

GitHub Actions will automatically check submitted changes.

```text
Pull Request
     ↓
Test
     ↓
Build
     ↓
Pass ✓
     ↓
Ready for Review / Merge
```

---

## Project Status

**Status:** In Development

### Completion Checklist

* [ ] Project setup
* [ ] Registration
* [ ] Login
* [ ] Authentication
* [ ] Session management
* [ ] Dashboard
* [ ] Logout
* [ ] Testing
* [ ] CI
* [ ] Code Review
* [ ] Final Integration
* [ ] Main Release

---

## Final Deliverable

A working Go authentication web application demonstrating both:

**Application Functionality**

and

**Team Software Development Workflow**

```text
User
 ↓
Register
 ↓
Login
 ↓
Session
 ↓
Dashboard
 ↓
Logout
```

---

## Repository

**Project:** `go-auth-demo`

**Development Model:** Feature Branch → Pull Request → Code Review → Develop → Main
