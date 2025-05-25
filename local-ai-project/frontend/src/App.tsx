import React, { useState, useEffect } from 'react';
import ModelManager from './components/ModelManager';
import DocumentUploader from './components/DocumentUploader';
import ChatInterface from './components/ChatInterface';
import WikiResults from './components/WikiResults';
import DocumentList from './components/DocumentList';
import { ApiService } from './services/api';
import type { Model } from './types/Model';
import type { Document } from './types/Document';
import type { QueryResponse } from './types/chatMessage';

const App: React.FC = () => {
    const [models, setModels] = useState<Model[]>([]);
    const [documents, setDocuments] = useState<Document[]>([]);
    const [selectedModel, setSelectedModel] = useState<string>('');
    const [currentResponse, setCurrentResponse] = useState<QueryResponse | null>(null);
    const [loading, setLoading] = useState(false);
    const [activeTab, setActiveTab] = useState<'chat' | 'models' | 'documents'>('chat');

    const apiService = new ApiService();

    useEffect(() => {
        loadModels();
        loadDocuments();
    }, []);

    const loadModels = async () => {
        try {
            const response = await apiService.getModels();
            setModels(response.models);
        } catch (error) {
            console.error('Failed to load models:', error);
        }
    };

    const loadDocuments = async () => {
        try {
            const response = await apiService.getDocuments();
            setDocuments(response.documents);
        } catch (error) {
            console.error('Failed to load documents:', error);
        }
    };

    const handleQuery = async (query: string, includeWiki: boolean = true) => {
        setLoading(true);
        try {
            const response = await apiService.query({
                query,
                include_wiki: includeWiki,
                model_name: selectedModel
            });
            setCurrentResponse(response);
        } catch (error) {
            console.error('Query failed:', error);
        } finally {
            setLoading(false);
        }
    };

    const handleModelDownload = async (name: string, url: string) => {
        try {
            await apiService.downloadModel(name, url);
            await loadModels();
        } catch (error) {
            console.error('Model download failed:', error);
        }
    };

    const handleDocumentUpload = async (file: File) => {
        try {
            await apiService.uploadDocument(file);
            await loadDocuments();
        } catch (error) {
            console.error('Document upload failed:', error);
        }
    };

    const handleModelLoad = async (modelName: string) => {
        try {
            await apiService.loadModel(modelName);
            setSelectedModel(modelName);
        } catch (error) {
            console.error('Model load failed:', error);
        }
    };

    return (
        <div className="App">
            <header className="app-header">
                <h1>🤖 Local AI Assistant</h1>
                <div className="status-bar">
                    <span className="model-status">
                        Model: {selectedModel || 'None selected'}
                    </span>
                    <span className="document-count">
                        Documents: {documents.length}
                    </span>
                </div>
            </header>

            <nav className="tab-navigation">
                <button
                    className={`tab ${activeTab === 'chat' ? 'active' : ''}`}
                    onClick={() => setActiveTab('chat')}
                >
                    💬 Chat
                </button>
                <button
                    className={`tab ${activeTab === 'models' ? 'active' : ''}`}
                    onClick={() => setActiveTab('models')}
                >
                    🎯 Models
                </button>
                <button
                    className={`tab ${activeTab === 'documents' ? 'active' : ''}`}
                    onClick={() => setActiveTab('documents')}
                >
                    📁 Documents
                </button>
            </nav>

            <main className="app-main">
                {activeTab === 'chat' && (
                    <div className="chat-section">
                        <div className="chat-container">
                            <ChatInterface
                                onQuery={handleQuery}
                                loading={loading}
                                response={currentResponse}
                            />
                        </div>
                        
                        {currentResponse && (
                            <div className="results-section">
                                <WikiResults 
                                    results={currentResponse.sources.wiki} 
                                    isLoading={false}
                                />
                            </div>
                        )}
                    </div>
                )}

                {activeTab === 'models' && (
                    <ModelManager
                        models={models}
                        selectedModel={selectedModel}
                        onDownload={handleModelDownload}
                        onLoad={handleModelLoad}
                        onRefresh={loadModels}
                    />
                )}

                {activeTab === 'documents' && (
                    <div className="documents-section">
                        <DocumentUploader onUpload={handleDocumentUpload} />
                        <DocumentList
                            documents={documents}
                            onRefresh={loadDocuments}
                        />
                    </div>
                )}
            </main>
        </div>
    );
};

export default App;