import axios from "axios";
import type { Model } from "../types/Model";
import type { Document, DocumentUploadResponse } from "../types/Document";
import type { WikiResult } from "../types/WikiResult";
import type { QueryResponse } from "../types/chatMessage";

const API_BASE_URL =
  import.meta.env.VITE_API_URL || "http://localhost:8082/api/v1";

const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 30000,
  headers: {
    "Content-Type": "application/json",
  },
});

export class ApiService {
  // Health check
  async healthCheck(): Promise<{ status: string; message: string }> {
    const response = await api.get("/health");
    return response.data;
  }

  // Model management
  async getModels(): Promise<{ models: Model[] }> {
    const response = await api.get("/models");
    return response.data;
  }

  async downloadModel(name: string, url: string): Promise<{ message: string }> {
    const response = await api.post("/models/download", { name, url });
    return response.data;
  }

  async loadModel(name: string): Promise<{ message: string }> {
    const response = await api.post("/models/load", { name });
    return response.data;
  }

  async deleteModel(name: string): Promise<{ message: string }> {
    const response = await api.delete(`/models/${name}`);
    return response.data;
  }

  // Document management
  async getDocuments(): Promise<{ documents: Document[] }> {
    const response = await api.get("/documents");
    return response.data;
  }

  async uploadDocument(file: File): Promise<DocumentUploadResponse> {
    const formData = new FormData();
    formData.append("file", file);

    const response = await api.post("/documents/upload", formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    });
    return response.data;
  }

  async deleteDocument(id: number): Promise<{ message: string }> {
    const response = await api.delete(`/documents/${id}`);
    return response.data;
  }

  // Wiki search
  async searchWiki(query: string): Promise<{ results: WikiResult[] }> {
    const response = await api.get(
      `/wiki/search?q=${encodeURIComponent(query)}`
    );
    return response.data;
  }

  // AI Query
  async query(params: {
    query: string;
    include_wiki: boolean;
    model_name: string;
  }): Promise<QueryResponse> {
    const response = await api.post("/query", {
      query: params.query,
      include_wiki: params.include_wiki,
      model_name: params.model_name,
    });
    return response.data;
  }
}

export default new ApiService();
