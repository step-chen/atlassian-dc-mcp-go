---
trigger: manual
name: cpp_architecture_design
---

# C++ Architecture Design

## Overview

This document contains architecture design specifications for C++ projects.

## **Technology Stack**
- **Language Version**: C++17
- **Build System**: CMake (3.15+)
- **Frameworks**: Qt 5.15 (for UI components), Boost 1.75+ (for utility functions)
- **Dependency Management**: Conan
- **Databases**: SQLite (embedded), PostgreSQL 12+ (remote)
- **Testing Tools**: Google Test, Google Mock
- **Build/Deployment**: Docker, Kubernetes
- **Static Analysis**: Clang Static Analyzer, Cppcheck
- **Documentation**: Doxygen

---

## **Application Logic Design**
### **Layered Design Specifications**
1. **Presentation Layer** (UI Components):
   - Handle user interactions and display data.
   - Use Qt for cross-platform UI implementation.
   - Depend on the Business Logic layer, **must not directly operate the database**.
   - Located in `src/ui/` directory with corresponding headers in `include/ui/`.
   
2. **Business Logic Layer** (Core Services):
   - Implement core business logic and call Data Access components.
   - Return results or errors, **do not directly handle UI events**.
   - Located in `src/core/` directory with corresponding headers in `include/core/`.
   
3. **Data Access Layer** (Repositories):
   - Encapsulate database operations (SQL queries).
   - Provide interface definitions and implement interaction with SQLite and PostgreSQL databases.
   - Located in `src/data/` directory with corresponding headers in `include/data/`.
   
4. **Entities Layer** (Domain Models):
   - Define domain objects (such as User, Product, Map Data).
   - **Do not contain business logic or database operations**.
   - Located in `include/models/` directory.
   
5. **Utilities Layer** (Helper Functions):
   - Encapsulate common functions (such as logging, file operations, string processing, math utilities).
   - Located in `src/utils/` directory with corresponding headers in `include/utils/`.

---

## **Generic C++ Project Architecture**

### **Module Structure**
C++ projects typically consist of several key modules:

1. **Data Parser Module**:
   - Parses specific data formats
   - Converts proprietary formats to internal representation
   - Handles different versions and variants of data formats

2. **Processing Engine**:
   - Implements core algorithms
   - Works with parsed data
   - Provides optimization features

3. **Visualization Module**:
   - Renders data using Qt
   - Provides interactive viewing capabilities
   - Supports different data layers and overlays

4. **Data Export/Import**:
   - Exports processed data to various formats
   - Imports external data for integration
   - Handles format conversions

### **Data Flow**
1. Raw data is read from files or databases
2. Parser module converts data to internal representation
3. Business logic processes the data as required
4. Visualization module displays the data to users
5. Results can be exported in various formats

### **Performance Considerations**
- Large datasets are handled with memory-mapped files
- Multi-threading is used for CPU-intensive operations
- Database connection pooling for remote PostgreSQL access
- Lazy loading for UI components to improve startup time