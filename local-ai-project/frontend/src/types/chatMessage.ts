export interface ChatMessage {
    id: string;
    role: "user" | "assistant" | "system";
    content: string;
    timestamp: string;
    sources?: Source[];
}

export interface QueryResponse {
    response: string;
    sources: {
        documents: Document[];
        wiki: WikiResult[];
    };
}

export interface Source {
    type: 'document' | 'wiki';
    title: string;
    content: string;
    relevanceScore: number;
    documentId?: string;
    url?: string;
}

// Import the types we need
import type { Document } from './Document';
import type { WikiResult } from './WikiResult';