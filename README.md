# 📌 Project Prompt: Go-Management POS System
## Goal:
> Build a Point of Sale system using Golang according to Clean Architecture, supporting product management, invoices, customers, inventory, import and export and reporting. 
>

![This is an alt text.](/docs/img/architect.png)
![This is an alt text.](/docs/img/image1.png)
![This is an alt text.](/docs/img/image.png)
# Project Structure (Clean Architecture)
```
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
```

# Main Use Case

## a. Auth & User
- [ ] Register / login with email, password
- [ ] Hash password (bcrypt)
- [ ] JWT Access & Refresh Token
- [ ] Role-based Access Control: admin, cashier, manager

## b. Product Management
- [ ] CRUD products: name, description, barcode, selling price, import price, inventory
- [ ] Product category

## c. Inventory Management
- [ ] Import management: supplier, import date, total amount
- [ ] Export management: by sales order
- [ ] Update inventory in real-time

## d. Sales & POS
- [ ] Create sales invoice: product, quantity, promotion, sales staff
- [ ] Calculate total amount, VAT, discount
- [ ] Print invoice (if needed, export PDF)

## e. Customer Management
- [ ] Add/edit/delete customers
- [ ] Track purchase history

## f. Reporting
- [ ] Revenue by day, week, month
- [ ] Top selling products
- [ ] Low inventory, need to add more


# Main Functional Modules
## a. Auth & User
- Features: Register/login, `JWT` tokens, role-based access (admin, cashier, manager).
- Implementation:
- Model:: User struct with fields like Email, PasswordHash, Role.
- Repository: Interface for CRUD operations on users.
- Service: Password hashing with bcrypt, `JWT` generation (access + refresh tokens).
- Delivery:: HTTP endpoints (/register, /login) using Fiber.
- Tech: `golang-jwt`/`jwt`, golang.org/x/crypto/bcrypt.
## b. Product Management
- Features: CRUD for products, categories.
- Implementation:
- Model:: Product (name, barcode, prices, stock), Category.
- Repository: Interface for product CRUD.
- Service: Business rules (e.g., validate barcode uniqueness).
- Delivery:: HTTP endpoints (/products, /products/{id}).
- Tech: Barcode generation via a library like github.com/boombuler/barcode.
## c. Inventory Management
- Features: Import/export tracking, real-time stock updates.
- Implementation:
- Model:: Import (supplier, date, items), Export (sales order link).
- Repository: Interface for inventory operations.
- Service: Update stock levels transactionally.
- Delivery:: HTTP endpoints (/inventory/import, /inventory/export).
- Tech: PostgreSQL for persistence, Redis for caching stock levels.
## d. Sales & POS
- Features: Invoice creation, calculations, PDF export.
- Implementation:
- Model:: Invoice (products, quantities, totals, staff).
- Repository: Interface for invoice CRUD.
- Service: Calculate totals (VAT, discounts), queue PDF generation.
- Delivery:: HTTP endpoint (/sales/invoice).
- Tech: `RabbitMQ` for async PDF generation, github.com/jung-kurt/gofpdf for PDF export.
## e. Customer Management
- Features: CRUD for customers, purchase history.
- Implementation:
- Model:: Customer (name, contact), PurchaseHistory.
- Repository: Interface for customer operations.
- Service: Link customers to invoices.
- Delivery:: HTTP endpoints (/customers, /customers/{id}/history).
## f. Reporting
- Features: Revenue, top products, low inventory.
- Implementation:
- Service: Aggregate data (e.g., SQL queries for revenue).
- Delivery:: HTTP endpoints (/reports/revenue, /reports/top-products).
- Tech: Redis for caching reports, PostgreSQL for raw data.
# Technology Stack 🚀
- Backend: `Golang` 🐹 + `Fiber` ⚡ (lightweight, fast HTTP framework).
- Database: `PostgreSQL` 🐘 with `GORM` 🛠️ (ORM) or SQLC 📜 (SQL-first approach for performance).
- Queue: `RabbitMQ` 🐰 for async tasks (e.g., printing invoices 🧾, sending emails 📧).
- Cache: Redis 💾 for caching products, reports.
- Auth: JWT 🔑 for tokens, bcrypt 🔒 for password hashing.
- Config: Viper 🐍 with `config.yaml` ⚙️ for environment-specific settings.
- Containerization: `Docker` 🐳 + `Docker Compose` 📦 for local dev and deployment.
- Testing: Go’s built-in testing package 🧪 for unit tests on the service layer.
- gRPC: Optional for inter-service communication 📡 (e.g., reporting service).