# XYZ Multifinance Backend Service

This project is a backend service built using Golang to support customer financing transactions for PT XYZ Multifinance. The service is designed to be scalable, secure, maintainable, and reliable, following Clean Architecture principles.

---

## Architecture

This project adopts **Clean Architecture** to ensure separation of concerns and maintainability.

```
Delivery Layer (HTTP / Fiber)
        ↓
Usecase Layer (Business Logic)
        ↓
Repository Layer (Data Access)
        ↓
Infrastructure Layer (Database, ORM, Config)
```

### Benefits
- High testability
- Easy feature extension
- Independent infrastructure replacement
- Clear business rule separation

---

## Project Structure

```
internal/
 ├── delivery      → HTTP handlers and routing
 ├── usecase       → Business logic
 ├── repository    → Database interaction
 ├── entity        → Domain models
 └── infrastructure→ Database & external services
```

---

## Database Design

### Tables

#### Customers
Stores customer personal information.

#### Customer Limits
Stores financing limit per tenor.

#### Transactions
Stores financing transactions performed by customers.

---

### Database Characteristics

- MySQL InnoDB engine
- ACID compliant
- Foreign key constraints
- Indexed queries for performance

---

## Concurrency Handling (Financial Safety)

Financial systems must prevent overspending when multiple transactions occur simultaneously.

This project ensures concurrency safety using:

```
SELECT ... FOR UPDATE
```

Implemented via GORM locking clause:

```go
tx.Clauses(clause.Locking{Strength: "UPDATE"})
```

### How It Works

1. Customer limit row is locked during transaction.
2. System validates remaining limit.
3. Limit usage is updated.
4. Transaction record is created.
5. Transaction is committed atomically.

### Result
- Prevents double spending
- Guarantees data consistency
- Supports high concurrent load

---

## Security (OWASP Top 10 Mitigation)

This service implements multiple security protections:

### 1. SQL Injection Prevention
- Using GORM prepared statements
- Parameterized queries

### 2. Input Validation
- Request payload validation
- Prevents malformed or malicious data

### 3. Rate Limiting
- Fiber middleware limiter
- Prevents abuse and brute force attacks

---

## Testing

Unit tests are implemented at the usecase layer.

Testing covers:

- Business rule validation
- Limit calculation logic
- Error scenario handling

---

## Docker Support

Application is containerized using Docker.

### Build
```
docker build -t xyz-multifinance .
```

### Run
```
docker run -p 8080:8080 xyz-multifinance
```

---

## Running Locally

### 1. Setup Environment
Create `config.json` file:

```
{
  "app": {
    "name": "xyz-multifinance"
  },
  "web": {
    "prefork": false,
    "port": 3000
  },
  "log": {
    "level": 6
  },
  "database": {
    "username": "xyz",
    "password": "password",
    "host": "localhost",
    "port": 3306,
    "name": "xyz_multifinance",
    "pool": {
      "idle": 10,
      "max": 100,
      "lifetime": 300
    }
  }
}
```

---

### 2. Run Migration
Execute:

```
migrations/schema.sql
migrations/seed.sql
```

---

### 3. Start Application
```
go run cmd/server/main.go
```

---

## 📡 API Example

### Create Transaction

```
POST /transactions
```

Request:
```json
{
  "customer_id": 1,
  "contract_number": "TRX-001",
  "otr": 100000,
  "admin_fee": 5000,
  "installment_amount": 20000,
  "interest_amount": 5000,
  "asset_name": "Motorcycle",
  "tenor_month": 3
}
```

---

## Scalability & Reliability

- Stateless REST API
- Horizontal scaling ready
- Database transaction isolation
- Container deployment ready
- Supports 99.9% availability target

---

## Seed Data

Includes sample customers:

- Budi Santoso
- Annisa Putri

Each customer has limits for tenor:
- 1 month
- 2 months
- 3 months
- 6 months

---

## Technology Stack

- Golang
- Fiber Web Framework
- GORM ORM
- MySQL
- Docker
- Testify (Unit Testing)

---

## Git Flow

Development follows Git Flow strategy:

- main → Production ready
- develop → Feature development

---

## Future Improvements

- Authentication & Authorization
- Distributed caching
- Message queue integration
- Monitoring & logging enhancement

---
