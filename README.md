/go-management-pos
├── /cmd              # Entry points for the application
│   ├── /pos          # Main POS application
│   └── /migrate     # Database migration tool
├── /internal         # Core application logic (private to the app)
│   ├── /domain       # Business entities and logic
│   │   ├── /model    # Structs (e.g., Product, User, Invoice)
│   │   ├── /repository # Interfaces for data access
│   │   ├── /service  # Business logic implementation
│   │   └── /usecase  # Use case layer (orchestrates services)
│   ├── /infrastructure # External systems (DB, cache, queue)
│   │   ├── /db       # PostgreSQL setup (GORM or SQLC)
│   │   ├── /redis    # Redis client
│   │   └── /queue    # RabbitMQ client
│   └── /delivery     # Handlers for external interaction
│       ├── /http     # REST API handlers (Fiber)
│       └── /grpc     # gRPC handlers (optional)
├── /pkg              # Shared utilities
│   ├── /logger       # Logging (e.g., Zap or Logrus)
│   ├── /validator    # Input validation
│   └── /helper       # Common utilities (e.g., barcode generation)
├── /config           # Configuration (Viper + YAML)
├── /docs             # API docs (e.g., Swagger)
├── /test             # Unit tests
├── Dockerfile        # Docker configuration
├── docker-compose.yml # Multi-container setup
└── go.mod            # Go module file
