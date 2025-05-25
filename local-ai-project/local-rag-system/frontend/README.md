# Frontend Project Documentation

## Overview

This project is a Retrieval-Augmented Generation (RAG) system that allows users to upload documents, manage AI models, and interact with a chat interface. The frontend is built using TypeScript and React, providing a user-friendly interface for interacting with the backend services.

## Project Structure

The frontend project is organized as follows:

- **src/**: Contains the source code for the application.
  - **components/**: Reusable React components.
    - `ChatInterface.tsx`: Handles user input and displays chat messages.
    - `DocumentUploader.tsx`: Allows users to upload documents for processing.
    - `ModelManager.tsx`: Manages AI model selection and operations.
    - `WikiResults.tsx`: Displays search results from the Wikipedia API.
  - **services/**: Contains API service functions.
    - `api.ts`: Functions for making API calls to the backend.
    - `types.ts`: TypeScript types and interfaces used throughout the application.
  - **store/**: State management setup.
    - `index.ts`: Configures state management using Zustand or Redux Toolkit.
  - `App.tsx`: Main application component that renders core components.
  - `main.tsx`: Entry point for the frontend application.
  - `index.css`: Global styles for the application.

- **public/**: Contains static assets.
  - `vite.svg`: SVG image used in the application.

- **package.json**: Configuration file for npm, listing dependencies and scripts.

- **tsconfig.json**: TypeScript configuration file.

- **vite.config.ts**: Vite configuration file.

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd local-rag-system/frontend
   ```

2. **Install dependencies:**
   ```
   npm install
   ```

3. **Run the development server:**
   ```
   npm run dev
   ```

4. **Open your browser and navigate to:**
   ```
   http://localhost:3000
   ```

## Usage

- Use the **Document Uploader** to upload documents for processing.
- Manage AI models using the **Model Manager** component.
- Interact with the chat interface to ask questions and receive responses.
- View results from the Wikipedia API in the **Wiki Results** component.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.