# 🚀 GoGooli Framework

**Tired of building Go APIs that turn into spaghetti? GoGooli is the clean slate you've been looking for.**

GoGooli is not another heavy, all-in-one web framework. Instead, it's a modern, production-ready **boilerplate** built on the principles of **Clean Architecture**. It provides a solid foundation and a clear, decoupled structure to build powerful, scalable, and maintainable backend services in Go.

---

## ✨ Core Philosophy: Building a High-End PC

We believe building a backend service is like assembling a custom gaming PC. You don't build the CPU or the RAM yourself. You choose the best components from the market and assemble them on a powerful motherboard that has standard slots.

**GoGooli is that powerful motherboard.**

Our architecture is designed to be completely decoupled:
* **The CPU (Core Logic):** Your application's business logic (`/internal/usecase`) is the brain. It's totally independent and knows nothing about your database or web framework.
* **The Motherboard Slots (Interfaces):** We define standard contracts (`/internal/repository`) for external parts like data storage. The CPU plugs into these slots, not directly into the hardware.
* **The Hardware (Implementations):** These are the concrete pieces (`/platform/database`). We chose PostgreSQL and `sqlc` as our high-performance "RAM," but you could swap it out for something else without changing the CPU.
* **You, The Builder (`main.go`):** The `main.go` file is where you, the architect, assemble all these high-end components together using Dependency Injection.

---

## 📋 Features At a Glance

* **Clean Architecture:** A strict separation of concerns that makes the code testable and maintainable.
* **High-Performance Web Server:** Built on [Gin](https://github.com/gin-gonic/gin) for fast and reliable HTTP routing.
* **Type-Safe Database Queries:** Uses [sqlc](https://github.com/sqlc-dev/sqlc) to generate fully type-safe Go code from your raw SQL. No magic ORMs, full control, and top performance.
* **Schema Migrations:** Integrated with [golang-migrate](https://github.com/golang-migrate/migrate) to manage your database schema's evolution.
* **Externalized Configuration:** Uses [Viper](https://github.com/spf13/viper) to load all settings from a `config.yml` file, keeping secrets out of your code.
* **Ready for Production:** With built-in support for Docker, getting from development to production is straightforward.

---

## 🚀 Getting Started

Follow these steps to get your GoGooli project up and running.

### 1. Prerequisites

* [Go](https://golang.org/dl/) (version 1.18 or higher)
* [Docker](https://www.docker.com/products/docker-desktop/)
* [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
* [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html)

### 2. Setup

1.  **Clone the Repository**
    
    git clone [https://github.com/raheemtermeh/GoGooli.git](https://github.com/raheemtermeh/GoGooli.git)
    cd GoGooli
   

2.  **Create Configuration File**
    Create a `config.yml` file in the root of the project and paste the following content. This file will hold all your application's settings.

    **`config.yml`**
    ```yaml
    database:
      url: "postgres://termeh:hemret@localhost:5432/goyandb?sslmode=disable"

    server:
      port: ":8080"
    ```

### 3. Launch the Database

We use Docker to run a PostgreSQL database. This command starts a clean database instance with the credentials defined in your `config.yml`.


docker run --name gogooli-db -e POSTGRES_USER=termeh -e POSTGRES_PASSWORD=hemret -e POSTGRES_DB=goyandb -p 127.0.0.1:5432:5432 -d postgres
4. Run Database Migrations
With the database running, apply the schema to create your tables. Wait a few seconds after starting the container for the database to be ready.


migrate -path migrations -database 'postgres://termeh:hemret@localhost:5432/goyandb?sslmode=disable' up
Tip: You can add sample data by connecting to the database with a tool like DBeaver and running INSERT queries.

5. Run the Application
You're all set! Start the server with this command:

go run cmd/api/main.go
You should see logs indicating the server has started successfully on port :8080.

6. Test Your API
Open a new terminal and use curl to test the products endpoint:


curl http://localhost:8080/api/products
🏛️ In-Depth: Architecture Explained
The structure of GoGooli is designed for maximum clarity and separation of concerns.

/cmd/api/main.go: The entry point. Its only job is to be the "assembler." It reads the config, initializes all the dependencies (database, repositories, use cases), "injects" them into each other, and starts the server. It represents the outermost layer of the application.

/internal/domain: The heart of your application. It contains the core data structures and business models (like Product) that are central to your business. This layer has zero external dependencies.

/internal/usecase: This layer contains the specific business rules and orchestrates the flow of data. It calls the repository interfaces to get data but knows nothing about how the data is actually stored.

/internal/repository: Defines the contracts (Go interfaces) for our data layer. It answers the question "What can we do with our data?" (e.g., GetAll()).

/platform: This directory holds all the code that interacts with the outside world.

/database: The concrete implementation of our repository interfaces. It knows how to talk to PostgreSQL using the code generated by sqlc.

/web: Contains our Gin server, routes, and HTTP handlers. It's responsible for translating HTTP requests into calls to our use cases.

This structure ensures that our core business logic (domain and usecase) remains pure and untouched by the details of our infrastructure (like the database or web framework).

🗺️ Future Roadmap
GoGooli is a living project. Here are some of the features planned for the future:

[ ] JWT Authentication: Middleware for securing endpoints.

[ ] Structured Logging: Integration with a logger like Zap or Logrus.

[ ] Request Validation: Add automatic request body validation.

[ ] CLI Tool: A powerful CLI to generate new modules (handlers, usecases, etc.) automatically.

[ ] Enhanced Observability: Add metrics (Prometheus) and tracing (OpenTelemetry).

🙌 Contributing
Contributions are welcome! If you have ideas for new features or improvements, feel free to open an issue or submit a pull request.

📜 License
This project is licensed under the MIT License.
