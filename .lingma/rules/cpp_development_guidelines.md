---
trigger: manual
name: cpp_development_guidelines
---

# C++ Development Guidelines

## Overview

This document contains development guidelines and coding standards for C++ projects.

## **Specific Development Specifications**

### **1. Package/Module Management**
- **Namespace Naming**:
  - Namespaces should be lowercase with clear structure (e.g. `project::module::component`).
  - Avoid circular dependencies between modules, use forward declarations when possible.
  - All code should be within a project-specific namespace.
- **Modularity**:
  - Each functional module should be an independent library (e.g. `parser`, `engine`, `visualization`, `utils`).
  - Libraries should have well-defined interfaces with minimal dependencies.

### **2. Code Structure**
- **File Organization**:
  ```
  project-root/
  ├── src/          
  │   ├── parser/     # Data parsing implementation
  │   ├── engine/     # Core algorithms
  │   ├── ui/         # Qt-based UI components
  │   ├── data/       # Database access implementations
  │   └── utils/      # Utility functions
  ├── include/      
  │   ├── parser/     # Headers for parser module
  │   ├── engine/     # Headers for engine module
  │   ├── ui/         # Headers for UI components
  │   ├── data/       # Headers for data access
  │   ├── models/     # Domain model definitions
  │   └── utils/      # Headers for utility functions
  ├── tests/          # Unit and integration tests
  ├── docs/           # Documentation files
  ├── CMakeLists.txt  # Main build configuration
  ├── conanfile.txt   # Dependencies
  └── README.md       # Project overview
  ```
- **Class Design**:
  - Classes should have a single responsibility with well-defined interfaces.
  - Follow the Rule of Five (destructor, copy constructor, copy assignment, move constructor, move assignment).
  - Use RAII for resource management.
  - Prefer composition over inheritance.
  - Mark classes as `final` when inheritance is not intended.

### **3. Error Handling**
- **Exception Safety**:
  ```cpp
  try {
      perform_operation();
  } catch (const std::exception& e) {
      log_error(e.what());
      throw; // Re-throw if needed
  }
  ```
- **Error Codes**:
  ```cpp
  enum class ErrorCode {
      Success = 0,
      InvalidInput = 1,
      FileNotFound = 2,
      ParseError = 3,
      DatabaseError = 4
  };
  
  struct Result {
      ErrorCode error_code;
      std::string message;
  };
  ```
- **Qt Error Handling**:
  - Use `Q_OBJECT` macro for classes that use signals/slots
  - Handle Qt-specific exceptions with `QException`

### **4. Memory Management**
- **Smart Pointers**:
  ```cpp
  // Use unique_ptr for exclusive ownership
  std::unique_ptr<MyClass> obj = std::make_unique<MyClass>();
  
  // Use shared_ptr for shared ownership
  std::shared_ptr<MyClass> shared_obj = std::make_shared<MyClass>();
  
  // Use weak_ptr to break circular references
  std::weak_ptr<MyClass> weak_obj = shared_obj;
  ```
- **STL Containers**:
  - Prefer `std::vector`, `std::string`, etc. over raw arrays.
  - Use `std::string_view` for non-owning string references (C++17).
  - Use `std::optional` for values that may or may not exist.

### **5. Modern C++ Features**
- **Auto Keyword**:
  ```cpp
  auto value = getSomeValue(); // Type inferred
  ```
- **Range-based For Loops**:
  ```cpp
  for (const auto& item : container) {
      process(item);
  }
  ```
- **Constexpr and Const Correctness**:
  - Use `constexpr` for compile-time constants
  - Mark functions as `const` when they don't modify object state
  - Use `const&` for function parameters that are not modified

### **6. Database Operations**
- **SQLite Usage**:
  ```cpp
  #include <sqlite3.h>
  
  class DatabaseManager {
  public:
      bool executeQuery(const std::string& query);
      std::vector<Record> fetchRecords(const std::string& query);
  private:
      sqlite3* db_;
  };
  ```
- **Connection Management**:
  - Use RAII for database connections
  - Implement connection pooling for PostgreSQL
  - Always sanitize inputs to prevent SQL injection

### **7. Concurrency Handling**
- **Thread Safety**:
  ```cpp
  #include <mutex>
  #include <thread>
  
  class ThreadSafeCounter {
  private:
      mutable std::mutex mtx;  // mutable for const methods
      int count = 0;
  
  public:
      void increment() {
          std::lock_guard<std::mutex> lock(mtx);
          ++count;
      }
      
      int getValue() const {
          std::lock_guard<std::mutex> lock(mtx);
          return count;
      }
  };
  ```
- **Qt Concurrency**:
  - Use `QtConcurrent` for parallel algorithms
  - Use `QThread` for long-running operations
  - Prefer signals/slots over direct function calls between threads

### **8. Security Specifications**
- **Input Validation**:
  ```cpp
  bool validateInput(const std::string& input) {
      if (input.empty() || input.size() > MAX_LENGTH) {
          return false;
      }
      // Additional validation logic
      return true;
  }
  ```
- **File Operations**:
  - Validate file paths to prevent directory traversal
  - Check file permissions before accessing
  - Sanitize user-provided file names

### **9. Testing Specifications**
- **Unit Testing with Google Test**:
  ```cpp
  #include <gtest/gtest.h>
  
  TEST(ParserTest, TestFileParsing) {
      // Arrange
      DataParser parser;
      std::string test_file = "test_data/sample.data";
      
      // Act
      auto result = parser.parseFile(test_file);
      
      // Assert
      EXPECT_TRUE(result.has_value());
      EXPECT_GT(result->getNodes().size(), 0);
  }
  ```
- **Test Organization**:
  - Tests should mirror the source structure in the `tests/` directory
  - Each module should have its own test subdirectory
  - Use test fixtures for shared setup/teardown

### **10. Logging Specifications**
- **Logging Framework**:
  ```cpp
  #include <spdlog/spdlog.h>
  
  void logInfo(const std::string& message) {
      spdlog::info("[Project Name] {}", message);
  }
  
  void logError(const std::string& message) {
      spdlog::error("[Project Name] {}", message);
  }
  ```
- **Log Levels**:
  - DEBUG: Detailed information for diagnosing problems
  - INFO: General information about program execution
  - WARN: Warning messages about potential issues
  - ERROR: Error events that might still allow the application to continue
  - CRITICAL: Serious errors that may lead to termination

### **11. Generic C++ Project Guidelines**
- **Data Handling**:
  - All data coordinates should use a consistent coordinate system
  - Implement proper coordinate transformation functions
  - Handle different data versions gracefully

- **Performance Optimization**:
  - Profile code regularly with Valgrind or similar tools
  - Minimize memory allocations in performance-critical paths
  - Use appropriate data structures for the data type (consider spatial indexing)

- **Qt UI Guidelines**:
  - Separate UI logic from business logic
  - Use Model-View architecture where appropriate
  - Implement proper signal/slot connections
  - Handle UI updates on the main thread

### **12. Code Documentation**
- **Doxygen Comments**:
  ```cpp
  /**
   * @brief Parses data files
   * @param filename Path to the data file
   * @return Parsed data or error information
   * @throws std::runtime_error if file cannot be opened
   */
  std::optional<Data> parseDataFile(const std::string& filename);
  ```
- **README Documentation**:
  - Each module should have a README.md explaining its purpose
  - Document build requirements and dependencies
  - Include usage examples for major components