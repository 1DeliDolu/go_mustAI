# Local RAG System

## Overview

The Local RAG (Retrieval-Augmented Generation) System is designed to facilitate the management and processing of documents while integrating AI models for enhanced query responses. This project consists of a backend service built with Go and a frontend application developed using TypeScript.

## Project Structure

The project is organized into two main directories: `backend` and `frontend`.

### Backend

The backend is responsible for handling API requests, processing documents, managing AI models, and integrating with external services like Wikipedia. The key components include:

- **cmd/server/main.go**: Entry point for the backend application.
- **internal/handlers/**: Contains functions for handling various API requests.
- **internal/models/**: Defines data models used in the application.
- **internal/services/**: Implements business logic for document processing, model management, and interactions with external APIs.
- **internal/storage/**: Manages database connections and vector storage for embeddings.
- **pkg/**: Contains configuration and utility functions.
- **api/openapi.yaml**: API specification in OpenAPI format.

### Frontend

The frontend provides a user interface for interacting with the backend services. It includes components for managing models, uploading documents, and displaying query results. Key components include:

- **src/components/**: Contains React components for the application.
- **src/services/**: Handles API calls to the backend.
- **src/store/**: Manages application state.
- **src/App.tsx**: Main application component.
- **src/main.tsx**: Entry point for the frontend application.

## Setup Instructions

### Backend

1. Navigate to the `backend` directory.
2. Run `go mod tidy` to install dependencies.
3. Start the server using `go run cmd/server/main.go`.
4. The API will be available at `http://localhost:8082/api/v1`.

### Frontend

1. Navigate to the `frontend` directory.
2. Run `npm install` to install dependencies.
3. Start the development server using `npm run dev`.
4. Access the application at `http://localhost:3000`.

## Usage

- Use the Model Manager to select and manage AI models.
- Upload documents using the Document Uploader.
- Interact with the AI model through the Chat Interface.
- View search results from Wikipedia in the Wiki Results component.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.
