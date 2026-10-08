# Go API Gateway

A Go-based API Gateway designed to act as the central entry point for backend services.

The project demonstrates API gateway concepts including **reverse proxy routing, JWT authorization, rate limiting, CORS handling, request logging, token management, background workers, scheduling, and service integration**.

## Architecture

```text
                         Clients
                            |
                            v
                  +-------------------+
                  |    API Gateway    |
                  +---------+---------+
                            |
          +-----------------+------------------+
          |                 |                  |
          v                 v                  v
      Middleware       Token Manager      Rate Limiter
          |                 |                  |
          +-----------------+------------------+
                            |
                            v
                    Reverse Proxy
                            |
          +-----------------+------------------+
          |                 |                  |
          v                 v                  v
      Backend A         Backend B          Backend C
```

The API Gateway provides a centralized layer for handling authentication, authorization, routing, traffic control, logging, and cross-origin requests.

## Project Structure

```text
.
├── bg_management
│   ├── processors
│   ├── scheduler
│   └── worker
│       └── worker.go
│
├── config
│   └── server.go
│
├── libvirt_exporter
│
├── middlewares
│   ├── cors_middleware.go
│   ├── entry_middlewares.go
│   ├── jwt_authorization.go
│   ├── logging.go
│   └── ratelimiter.go
│
├── routes
│   └── routes.go
│
├── routing
│   └── reverse_pro.go
│
├── token_manager
│   ├── create.go
│   ├── create.ho
│   ├── handler.go
│   └── validator.go
│
├── main.go
├── go.mod
└── go.sum
```

## Core Components

### API Gateway

`main.go`

The application entry point.

It initializes the gateway, configuration, routes, middleware, and supporting components.

---

## Routing

### `routes/routes.go`

Contains API route definitions.

Routes determine which requests are handled by the gateway and which backend service should receive the request.

Typical flow:

```text
Client Request
      |
      v
   Router
      |
      v
 Middleware
      |
      v
 Reverse Proxy
      |
      v
 Backend Service
```

### `routing/reverse_pro.go`

Contains reverse proxy functionality.

The gateway receives a client request and forwards it to the appropriate backend service.

```text
Client
  |
  | HTTP Request
  v
API Gateway
  |
  | Forward Request
  v
Backend Service
  |
  | Response
  v
API Gateway
  |
  v
Client
```

This allows backend services to remain hidden behind a single public gateway.

## Middleware

The project contains multiple middleware components.

### CORS

`middlewares/cors_middleware.go`

Handles **Cross-Origin Resource Sharing (CORS)**.

It controls which external origins are allowed to communicate with the API Gateway.

### JWT Authorization

`middlewares/jwt_authorization.go`

Provides JWT-based authorization.

Typical flow:

```text
Client
  |
  | JWT
  v
API Gateway
  |
  v
JWT Validation
  |
  +---- Invalid ---> Unauthorized
  |
  +---- Valid -----> Backend
```

### Logging

`middlewares/logging.go`

Provides request/response logging.

Useful information can include:

- HTTP method
- Request path
- Response status
- Request duration
- Errors
- Client information

### Rate Limiting

`middlewares/ratelimiter.go`

Controls how many requests a client can make within a given period.

```text
Client
  |
  v
Rate Limiter
  |
  +---- Limit exceeded ---> HTTP 429
  |
  +---- Allowed ----------> Backend
```

Rate limiting helps protect backend services from excessive traffic and abuse.

### Entry Middleware

`middlewares/entry_middlewares.go`

Provides common middleware initialization and request-processing logic used when requests enter the gateway.

## Token Management

The `token_manager` package handles token-related operations.

```text
token_manager/
├── create.go
├── create.ho
├── handler.go
└── validator.go
```

### Token Creation

`create.go`

Responsible for creating authentication tokens.

### Token Handler

`handler.go`

Handles token-related request processing.

### Token Validation

`validator.go`

Validates incoming tokens before protected requests are forwarded to backend services.

The general authentication flow is:

```text
             Client
                |
                v
         Authentication
                |
                v
          Token Manager
                |
                v
          JWT Validation
                |
          +-----+-----+
          |           |
        Valid       Invalid
          |           |
          v           v
      Backend      401/403
```

## Background Management

The `bg_management` package contains background processing components.

```text
bg_management/
├── processors/
├── scheduler/
└── worker/
    └── worker.go
```

These components can be used for tasks that should execute outside the main request/response lifecycle.

Typical architecture:

```text
API Request
    |
    v
Gateway
    |
    +---- Immediate Response
    |
    +---- Background Task
              |
              v
          Scheduler
              |
              v
           Worker
              |
              v
          Processor
```

This prevents long-running background operations from unnecessarily blocking API requests.

## Worker

`bg_management/worker/worker.go`

Contains worker-related functionality for processing background jobs.

Workers can be used for:

- Asynchronous processing
- Scheduled jobs
- Cleanup operations
- Resource management
- Long-running tasks

## Scheduler

The scheduler component is responsible for triggering background tasks based on configured schedules.

```text
Scheduler
    |
    +---- Job 1
    |
    +---- Job 2
    |
    +---- Job 3
```

## Processors

Processors contain the logic required to execute individual background tasks.

This separation helps keep scheduling and business logic independent.

## Libvirt Exporter

The `libvirt_exporter` component is intended for exposing or collecting **Libvirt-related metrics**.

This can be used to integrate infrastructure/resource metrics with monitoring systems.

Possible metrics include:

- CPU utilization
- Memory usage
- Domain status
- Network statistics
- Storage statistics
- Virtualization resource information

## Request Lifecycle

A typical request flows through the gateway as follows:

```text
                 Client
                   |
                   v
            +-------------+
            | API Gateway |
            +------+------+
                   |
                   v
              CORS Check
                   |
                   v
             Rate Limiter
                   |
                   v
           Request Logging
                   |
                   v
          JWT Authorization
                   |
                   v
               Router
                   |
                   v
           Reverse Proxy
                   |
                   v
            Backend Service
                   |
                   v
              Response
                   |
                   v
                Client
```

## Security

The gateway provides several security-related capabilities:

- JWT authentication/authorization
- Request validation
- Rate limiting
- CORS controls
- Centralized access control
- Request logging

Centralizing these controls allows backend services to focus primarily on business logic.

## Configuration

Configuration is maintained under:

```text
config/server.go
```

Configuration can include:

- Server port
- Backend service addresses
- JWT configuration
- Rate limiting configuration
- Proxy configuration
- Logging configuration

For production environments, sensitive configuration should be supplied through environment variables or a secrets-management system instead of hardcoding credentials.

## Requirements

- Go
- Backend services required by the configured routes
- Libvirt, if the exporter functionality is enabled

Check Go:

```bash
go version
```

## Installation

Clone the repository:

```bash
git clone <repository-url>
cd go-api-gateway
```

Download dependencies:

```bash
go mod download
```

## Run

Start the gateway:

```bash
go run .
```

Or:

```bash
go run main.go
```

## Build

Build the application:

```bash
go build -o api-gateway .
```

Run the binary:

```bash
./api-gateway
```

## Code Quality

Format the code:

```bash
go fmt ./...
```

Run static analysis:

```bash
go vet ./...
```

Run tests:

```bash
go test ./...
```

Run tests with the race detector:

```bash
go test -race ./...
```

## Go Concepts Demonstrated

This project can also be used to explore several important Go concepts:

- HTTP servers
- HTTP middleware
- Reverse proxy
- Goroutines
- Background workers
- Scheduling
- Concurrency
- Channels
- Context management
- Error handling
- JWT authentication
- Rate limiting
- Network programming
- Resource management
- Performance monitoring

## Design Goals

The main goals of this project are:

1. Provide a centralized API entry point.
2. Route requests to backend services.
3. Implement JWT-based authorization.
4. Apply rate limiting.
5. Handle CORS.
6. Provide centralized request logging.
7. Support background processing.
8. Provide infrastructure monitoring capabilities.
9. Hide internal backend service topology from clients.
10. Provide a foundation for scalable microservice architectures.

## Future Improvements

Possible improvements include:

- Distributed rate limiting using Redis
- Service discovery
- Circuit breaker integration
- Retry mechanisms
- Request tracing
- OpenTelemetry integration
- Prometheus metrics
- Health-check endpoints
- API versioning
- Graceful shutdown
- Configuration management
- TLS/mTLS support
- Request/response metrics
- Distributed background job queues

## License

This project is licensed under the terms specified in the `LICENSE` file.