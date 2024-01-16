# Go Counter Application

## Overview

This project is a simple REST-based Go application with a basic counter functionality. It allows users to create counters, increment them, retrieve the current value of a specific counter, and list all counters with their current values.

## Getting Started

### Prerequisites

- Go 1.18 or higher

### Installation

1. Clone the repository:

   ```
   git clone git@github.com:jashif/counters.git

   ```

2. Navigate to the project directory:
3. Download and install the project dependencies:

### Running the Application

Run the application using:

```bash
go run ./cmd/server/main.go
```

The server will start on port 8080.

## API Endpoints

- `POST /create`: Create a new counter (JSON payload: `{"name": "counterName"}`).
- `POST /increment`: Increment a named counter (Query parameter: `name=counterName`).
- `GET /value`: Get the current value of a counter (Query parameter: `name=counterName`).
- `GET /counters`: Get a list of all counters with their current values.

## Project Structure and Discussion

### 1. Project Structure

The project is organized into several packages:

- `cmd`: Contains the main application entry point.
- `counter`: Encapsulates the business logic and data access layer for counters.
- `handler`: Manages the HTTP request handlers, providing an interface to the application logic.
- `models`: Manages the domain and request models

### 2. Bootstrapping

The project is bootstrapped using Go modules. The `cmd/server/main.go` is the starting point, setting up dependencies and routing.

### 3. Testing Approach

- **Unit tests** for `counter` package to test business logic.
- **Mock tests** for `handler` package to test HTTP endpoints behavior.
- Integration tests to ensure the entire application functions correctly.

### 4. Database Choice

An in-memory data store is currently used for simplicity. For a production setup, a database like Redis or PostgreSQL could be considered for persistence and scalability.

### 5. Deployment Considerations

Containerization with Docker and Serverless can be used for deployment, ensuring consistency and scalability. CI/CD pipelines can automate testing and deployment processes.

Frontend can be hosted as static assests in S3 or cloud storage

### 6. Hosting

The application can be hosted on cloud platforms like AWS, GCP, or Azure. Managed services for databases and container orchestration can be leveraged for reliability and ease of maintenance.

## Contributions

Contributions, issues, and feature requests are welcome. Feel free to check [issues page](link-to-your-issue-page) if you want to contribute.

## License

Distributed under the MIT License. See `LICENSE` for more information.

Project Link: [https://github.com/jashif/counters](https://github.com/jashif/counters)

## Acknowledgments

- Gorilla Mux for the router functionality.
