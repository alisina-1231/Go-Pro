# Go-Pro

A REST API backend for a social feed app built in Go.

Users can register, log in, and view feeds. The project demonstrates Go best practices, HTTP handling, authentication, and modular code organization.

## Features

* User registration and authentication (JWT-based)
* Feed management: view and interact with posts
* Middleware for authentication and request handling
* Structured logging and error handling
* Modular design with `internal` packages for auth and database

## Prerequisites

* Go >= 1.20 installed
* Git
* Optional: Database (PostgreSQL/MySQL) configuration

## Installation

1. Clone the repository:

```bash
git clone https://github.com/alisina-1231/Go-Pro.git
cd Go-Pro
```

2. Initialize Go modules:

```bash
go mod tidy
```

3. Build the project:

```bash
go build
```

4. Run the project:

```bash
./Go-Pro
```

The server will run on port `8090` by default.

## Usage

Example endpoints:

* Register a new user:

```bash
POST /user/register
```

* Login:

```bash
POST /user/login
```

* Get user feed:

```bash
GET /feed
```

Logs will be printed to the console.

## Project Structure

```
Go-Pro/
├── main.go           # Entry point of the application
├── handler_user.go   # Handles user registration, login
├── handler_feed.go   # Handles feed endpoints
├── middleware_auth.go# Authentication middleware
├── internal/         # Internal packages (auth, database)
├── go.mod
└── go.sum
```

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Commit your changes: `git commit -m "Description of change"`
4. Push to the branch: `git push origin feature-name`
5. Open a Pull Request

## License

This project is licensed under the MIT License.
