# Local RAG System - Backend

This project is a Retrieval-Augmented Generation (RAG) system designed to facilitate document uploads and AI model management. The backend is built using Go and provides a RESTful API for interacting with various components of the system.

## Project Structure

The backend is organized into several key directories:

- **cmd/server**: Contains the entry point of the application.
- **internal**: Contains the core logic of the application, including handlers, models, services, and storage.
  - **handlers**: Functions for handling API requests.
  - **models**: Definitions of data structures used in the application.
  - **services**: Business logic for processing documents, managing models, and interacting with external APIs.
  - **storage**: Database and vector storage management.
- **pkg**: Contains utility functions and configuration management.
- **api**: OpenAPI specification for the API endpoints.

## Getting Started

### Prerequisites

- Go 1.16 or later
- A working database (SQLite or PostgreSQL)

### Installation

1. Clone the repository:

   ```
   git clone https://github.com/yourusername/local-rag-system.git
   cd local-rag-system/backend
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Configure your environment variables as needed.

### Running the Application

To start the backend server, run:

```
go run cmd/server/main.go
```

The server will start on `http://localhost:8082`.

### API Endpoints

- **GET /api/v1/models**: List available models.
- **POST /api/v1/models/download**: Download a specified model.
- **POST /api/v1/documents/upload**: Upload a document for processing.
- **POST /api/v1/query**: Submit a query to the AI model.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.
