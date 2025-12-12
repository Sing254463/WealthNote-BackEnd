# Go Fiber API

This project is a RESTful API built using Go Fiber, featuring live reloading with Go Air, JWT authentication, and support for both NoSQL (MongoDB) and SQL (MySQL and SQL Server) databases. The API is designed to be flexible and configurable through an `.env` file, allowing for easy adjustments without code modifications.

## Features

- **Live Reloading**: Utilizes Go Air for automatic reloading during development.
- **JWT Authentication**: Implements two types of tokens:
  - Short-lived token (15 minutes)
  - Long-lived token (15 days)
- **Database Support**: Connects to multiple databases:
  - MongoDB
  - MySQL
  - SQL Server
- **API Documentation**: Integrated with Go Swagger for easy API testing and documentation.
- **Configurable**: Uses an `.env` file for configuration, allowing you to enable or disable database connections and OAuth 2.0 features.

## Project Structure

```
WealthNoteBackend
├── cmd
│   └── main.go
├── internal
│   ├── config
│   │   └── config.go
│   ├── controllers
│   │   ├── auth.go
│   │   └── user.go
│   ├── middleware
│   │   ├── auth.go
│   │   └── cors.go
│   ├── models
│   │   └── user.go
│   ├── routes
│   │   └── routes.go
│   ├── services
│   │   ├── auth.go
│   │   └── user.go
│   └── database
│       ├── mongodb.go
│       ├── mysql.go
│       └── sqlserver.go
├── pkg
│   ├── jwt
│   │   └── jwt.go
│   └── utils
│       └── response.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── .air.toml
├── .env
├── .env.example
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Setup Instructions

1. **Clone the Repository**:
   ```
   git clone <repository-url>
   cd WealthNoteBackend
   ```

2. **Install Dependencies**:
   ```
   go mod tidy
   ```

3. **Configure Environment Variables**:
   - Copy `.env.example` to `.env` and update the values as needed.
   - Specify database connection settings and OAuth 2.0 options.

4. **Run the Application**:
   ```
   make run
   ```

5. **Access API Documentation**:
   - The Swagger UI can be accessed at `http://localhost:3000/docs`.

## Usage

- **Authentication**: Use the `/auth/login` endpoint to obtain JWT tokens.
- **User Management**: Access user-related endpoints for creating and retrieving user information.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.