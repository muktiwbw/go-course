# Blog API

A simple RESTful API for managing blog posts, built with Go and following Clean Architecture principles with proper abstractions.

## Project Structure

The project follows Clean Architecture principles with clear separation between interfaces and implementations:

```
.
├── .env              # Environment variables (not in version control)
├── .env.example      # Example environment template
├── pkg/              # Public interfaces and contracts
│   └── http/         # HTTP interfaces
│       ├── handler.go    # PostHandler interface
│       ├── router.go     # Router interface
│       └── server.go     # Server interface
├── internal/         # Private application code
│   ├── domain/       # Enterprise business rules
│   │   ├── errors.go # Domain-specific errors
│   │   └── post.go   # Post entity and interfaces
│   ├── usecase/      # Application business rules
│   │   └── post_usecase.go
│   ├── repository/   # Data storage implementation
│   │   └── post_repository.go
│   └── delivery/     # External interfaces implementations
│       └── http/
│           ├── handler/
│           │   └── post_handler.go   # PostHandler implementation
│           ├── router/
│           │   └── router.go         # Chi router implementation
│           └── server/
│               └── server.go         # HTTP server implementation
└── main.go          # Application entry point
```

## Architecture Overview

### Public Interfaces (`pkg/`)
Contains interfaces that can be implemented by other projects:

1. **HTTP Interfaces** (`pkg/http/`)
   - `PostHandler`: Defines HTTP endpoints for post management
   - `Router`: Defines routing contract
   - `Server`: Defines server behavior

### Private Implementation (`internal/`)

1. **Domain Layer** (`internal/domain/`)
   - Core business logic and rules
   - Domain entities and interfaces
   - No external dependencies

2. **Use Case Layer** (`internal/usecase/`)
   - Application-specific business rules
   - Orchestrates data flow using domain entities
   - Implements domain interfaces

3. **Repository Layer** (`internal/repository/`)
   - Data persistence implementation
   - Implements domain repository interfaces
   - Currently uses in-memory storage

4. **Delivery Layer** (`internal/delivery/http/`)
   - HTTP implementation details
   - Organized into:
     - `handler/`: Implements PostHandler interface
     - `router/`: Implements Router interface using Chi
     - `server/`: Implements Server interface

## Environment Configuration

The application uses environment variables for configuration:

```env
SERVER_PORT=8080  # HTTP server port
```

To get started:
1. Copy `.env.example` to `.env`
2. Adjust values as needed
3. Environment variables will be loaded at startup

## Clean Architecture Benefits

1. **Interface Segregation**
   - Clear separation between interfaces and implementations
   - Small, focused interfaces
   - Easy to understand and maintain

2. **Dependency Inversion**
   - High-level modules depend on abstractions
   - Implementations can be easily swapped
   - Better testability through interfaces

3. **Package Organization**
   - Public interfaces in `pkg/` for reuse
   - Private implementations in `internal/`
   - Clear boundaries between components

4. **Testability**
   - Interfaces make mocking easy
   - Components can be tested in isolation
   - No external dependencies needed for tests

5. **Flexibility**
   - Easy to replace implementations
   - Can switch router (e.g., from Chi to Gin)
   - Can change storage without affecting other parts

## API Endpoints

- `POST /posts` - Create a new post
- `GET /posts` - Get all posts
- `GET /posts/{id}` - Get a specific post
- `PUT /posts/{id}` - Update a post

## Getting Started

1. Clone the repository
2. Copy `.env.example` to `.env`
3. Run `go mod tidy` to install dependencies
4. Run `go run main.go` to start the server

The server will start on the port specified in your `.env` file.

## Future Improvements

1. Add database integration
2. Implement authentication
3. Add request validation
4. Add API documentation
5. Add metrics and monitoring
6. Implement caching
7. Add CI/CD pipeline 