---
trigger: manual
name: architecture_design
---

# Architecture Design

## Overview

This document contains architecture design specifications for the project.

## **Technology Stack**
- **Language Version**: Go 1.20+.
- **Frameworks**: Gin (HTTP framework), GORM (ORM), Zap (logging library).
- **Dependency Management**: Go Modules.
- **Databases**: PostgreSQL/MySQL (handwritten SQL or ORM).
- **Testing Tools**: Testify, Ginkgo.
- **Build/Deployment**: Docker, Kubernetes.

---

## **Application Logic Design**
### **Layered Design Specifications**
1. **Presentation Layer** (HTTP Handler):
   - Handle HTTP requests and convert request parameters to Use Cases.
   - Return structured JSON responses.
   - Depend on the Use Case layer, **must not directly operate the database**.
2. **Use Case Layer** (Business Logic):
   - Implement core business logic and call Repositories.
   - Return results or errors, **do not directly handle HTTP protocol**.
3. **Repository Layer** (Data Access):
   - Encapsulate database operations (such as GORM or handwritten SQL).
   - Provide interface definitions and implement interaction with specific databases.
4. **Entities Layer** (Domain Model):
   - Define domain objects (such as User, Product).
   - **Do not contain business logic or database operations**.
5. **DTOs Layer** (Data Transfer Objects):
   - Used for cross-layer data transfer (such as HTTP requests/responses).
   - Defined using `struct`, avoid duplication with Entities.
6. **Utilities Layer** (Utility Functions):
   - Encapsulate common functions (such as logging, encryption, time processing).