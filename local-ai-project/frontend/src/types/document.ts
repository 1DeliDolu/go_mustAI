export interface Document {
  id: string;
  name: string;
  type: 'pdf' | 'docx' | 'txt' | 'md';
  size: number;
  uploadDate: string;
  status: 'processing' | 'ready' | 'error';
  chunks?: number;
  embeddings?: boolean;
}

export interface DocumentUploadResponse {
  success: boolean;
  document?: Document;
  message: string;
}

export interface DocumentListProps {
  documents: Document[];
  onRefresh: () => Promise<void>;
}