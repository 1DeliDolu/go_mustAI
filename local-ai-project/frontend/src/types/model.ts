export interface Model {
  id: string;
  name: string;
  size: string;
  status: 'available' | 'downloading' | 'loading' | 'loaded' | 'error';
  downloadProgress?: number;
  description?: string;
  modelType: 'chat' | 'embedding' | 'multimodal';
}