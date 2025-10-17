# **Management Server API**

This project provides a management server for issuing and tracking Install App commands to agents, Gin for HTTP routing, PostgreSQL for database interactions, and Zap for logging.

## **Prerequisites**

Make sure the following are installed on your machine:

1. **Go** (version 1.20 or higher)
2. **PostgreSQL** (for database connectivity)

# 1. Install dependencies:

    go mod tidy

# 2. Set up PostgreSQL:

    Create a PostgreSQL database and user.
    Add the credentials to the .env file (see Configuration).

# 3. Create a .env file:

    Create a .env file at the root of the project with the necessary environment variables (see Configuration).

# 4. Configuration

    Create a .env file at the root of your project with the following variables:
            # Server Configuration
            SERVER_PORT=8180

    # Database Configuration
            DB_HOST=localhost
            DB_PORT=5432
            DB_USER=your-db-user
            DB_PASSWORD=your-db-password
            DB_NAME=your-db-name

# 5. Running the Project

    go run main.go

# 6. Access the API:

    The server will be available at http://localhost:8080 (or the port specified in the .env file).

# 7. Swagger UI

    http://localhost:8080/swagger-ui

# 8. Additional Information

    The Gin framework is used for HTTP request handling.
    Zap is used for structured logging with support for JSON and console formats.
