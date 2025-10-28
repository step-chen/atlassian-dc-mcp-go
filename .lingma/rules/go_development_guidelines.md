---
trigger: manual
name: development_guidelines
---

# Development Guidelines

## Overview

This document contains development guidelines and coding standards for the project.

## **Specific Development Specifications**

### **1. Package Management**
- **Package Naming**:
  - Package names should be lowercase with clear structure (e.g. `internal/repository`).
  - Avoid circular dependencies, use `go mod why` to check dependency relationships.
- **Modularity**:
  - Each function should be an independent sub-package (e.g. `cmd/api`, `internal/service`, `pkg/utils`).

### **2. Code Structure**
- **File Organization**:
  ```
  project-root/
  ├── cmd/          # Main entry (e.g. main.go)
  ├── internal/     # Core business logic
  │   ├── service/  # Business logic layer
  │   └── repository/ # Data access layer
  ├── pkg/          # Public utility packages
  ├── test/         # Test files
  └── go.mod        # Module dependencies
  ```
- **Function Design**:
  - Functions should have a single responsibility with no more than 5 parameters.
  - Use `return err` to explicitly return errors, **do not ignore errors**.
  - Defer resource release (e.g. `defer file.Close()`).

### **3. Error Handling**
- **Error Propagation**:
  ```go
  func DoSomething() error {
      if err := validate(); err != nil {
          return fmt.Errorf("validate failed: %w", err)
      }
      // ...
      return nil
  }
  ```
- **Custom Error Types**:
  ```go
  type MyError struct {
      Code    int    `json:"code"`
      Message string `json:"message"`
  }
  func (e *MyError) Error() string { return e.Message }
  ```

### **4. Dependency Injection**
- **Using Dependency Injection Frameworks**:
  ```go
  // Define interface
  type UserRepository interface {
      FindByID(ctx context.Context, id int) (*User, error)
  }
  
  // Implement dependency injection (e.g. using wire)
  func InitializeDependencies() (*UserRepository, func()) {
      repo := NewGORMUserRepository()
      return repo, func() { /* release resources */ }
  }
  ```

### **5. HTTP Handling**
- **Route Design**:
  ```go
  router := gin.Default()
  v1 := router.Group("/api/v1")
  {
      v1.POST("/users", CreateUserHandler)
      v1.GET("/users/:id", GetUserHandler)
  }
  ```
- **Response Format**:
  ```go
  type APIResponse struct {
      Status  string      `json:"status"`
      Message string      `json:"message"`
      Data    interface{} `json:"data,omitempty"`
  }
  ```

### **6. Database Operations**
- **GORM Usage Specifications**:
  ```go
  type User struct {
      gorm.Model
      Name  string `gorm:"unique"`
      Email string
  }
  
  func (repo *GORMUserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
      var user User
      if err := repo.DB.Where("email = ?", email).First(&user).Error; err != nil {
          return nil, err
      }
      return &user, nil
  }
  ```
- **SQL Injection Prevention**:
  - Use parameterized queries (e.g. `WHERE id = ?`).
  - Avoid concatenating SQL strings.

### **7. Concurrency Handling**
- **Goroutine Safety**:
  ```go
  var mu sync.Mutex
  var count int

  func Increment() {
      mu.Lock()
      defer mu.Unlock()
      count++
  }
  ```
- **Channel Communication**:
  ```go
  func Worker(id int, jobs <-chan int, results chan<- int) {
      for j := range jobs {
          fmt.Printf("Worker %d processing job %d\n", id, j)
          results <- j * 2
      }
  }
  ```

### **8. Security Specifications**
- **Input Validation**:
  ```go
  type CreateUserRequest struct {
      Name  string `json:"name" validate:"required,min=2"`
      Email string `json:"email" validate:"required,email"`
  }
  ```
- **Environment Variables**:
  ```go
  const (
      DBHost     = os.Getenv("DB_HOST")
      DBUser     = os.Getenv("DB_USER")
      DBPassword = os.Getenv("DB_PASSWORD")
  )
  ```

### **9. Testing Specifications**
- **Unit Testing**:
  ```go
  func TestUserService_CreateUser(t *testing.T) {
      // Use mock objects to simulate dependencies
      mockRepo := &MockUserRepository{}
      service := NewUserService(mockRepo)
      _, err := service.CreateUser(context.Background(), "test@example.com")
      assert.NoError(t, err)
  }
  ```

### **10. Logging Specifications**
- **Structured Logging**:
  ```go
  logger, _ := zap.NewProduction()
  defer logger.Sync()
  logger.Info("user created", zap.String("user_id", "123"))
  ```