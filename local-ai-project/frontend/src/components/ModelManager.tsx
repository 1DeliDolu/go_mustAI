// frontend/src/components/ModelManager.tsx
import React, { useState } from 'react';
import type { Model } from '../types/Model';

interface ModelManagerProps {
    models: Model[];
    selectedModel: string;
    onDownload: (name: string, url: string) => void;
    onLoad: (name: string) => void;
    onRefresh: () => void;
}

const ModelManager: React.FC<ModelManagerProps> = ({
    models,
    selectedModel,
    onDownload,
    onLoad,
    onRefresh
}) => {
    const [showDownloadForm, setShowDownloadForm] = useState(false);
    const [modelName, setModelName] = useState('');
    const [modelUrl, setModelUrl] = useState('');

    const handleDownload = (e: React.FormEvent) => {
        e.preventDefault();
        if (modelName && modelUrl) {
            onDownload(modelName, modelUrl);
            setModelName('');
            setModelUrl('');
            setShowDownloadForm(false);
        }
    };

    const popularModels = [
        { name: 'llama2', description: 'Meta Llama 2 - General purpose model' },
        { name: 'codellama', description: 'Code Llama - Specialized for coding' },
        { name: 'mistral', description: 'Mistral 7B - Fast and efficient' },
        { name: 'neural-chat', description: 'Intel Neural Chat - Conversational AI' }
    ];

    return (
        <div className="model-manager">
            <div className="section-header">
                <h2>🎯 AI Models</h2>
                <div className="header-actions">
                    <button onClick={onRefresh} className="refresh-button">
                        🔄 Refresh
                    </button>
                    <button
                        onClick={() => setShowDownloadForm(!showDownloadForm)}
                        className="add-button"
                    >
                        ➕ Add Model
                    </button>
                </div>
            </div>

            {showDownloadForm && (
                <div className="download-form">
                    <h3>Download New Model</h3>
                    <form onSubmit={handleDownload}>
                        <input
                            type="text"
                            placeholder="Model name (e.g., llama2)"
                            value={modelName}
                            onChange={(e) => setModelName(e.target.value)}
                            className="form-input"
                        />
                        <input
                            type="url"
                            placeholder="Model URL or Ollama model name"
                            value={modelUrl}
                            onChange={(e) => setModelUrl(e.target.value)}
                            className="form-input"
                        />
                        <div className="form-actions">
                            <button type="submit" className="download-button">
                                Download
                            </button>
                            <button
                                type="button"
                                onClick={() => setShowDownloadForm(false)}
                                className="cancel-button"
                            >
                                Cancel
                            </button>
                        </div>
                    </form>
                </div>
            )}

            <div className="popular-models">
                <h3>Popular Models</h3>
                <div className="model-grid">
                    {popularModels.map((model) => (
                        <div key={model.name} className="model-card popular">
                            <h4>{model.name}</h4>
                            <p>{model.description}</p>
                            <button
                                onClick={() => onDownload(model.name, model.name)}
                                className="download-button small"
                            >
                                Download via Ollama
                            </button>
                        </div>
                    ))}
                </div>
            </div>

            <div className="installed-models">
                <h3>Installed Models ({models.length})</h3>
                <div className="model-list">
                    {models.map((model) => (
                        <div
                            key={model.name}
                            className={`model-item ${selectedModel === model.name ? 'selected' : ''}`}
                        >
                            <div className="model-info">
                                <h4>{model.name}</h4>
                                <span className="model-size">
                                    {(typeof model.size === 'number' ? (model.size / 1024 / 1024).toFixed(1) : '?')} MB
                                </span>
                                <span className={`model-status ${model.status}`}>
                                    {model.status}
                                </span>
                            </div>
                            <div className="model-actions">
                                <button
                                    onClick={() => onLoad(model.name)}
                                    className={`load-button ${selectedModel === model.name ? 'active' : ''}`}
                                    disabled={selectedModel === model.name}
                                >
                                    {selectedModel === model.name ? '✅ Loaded' : '🔄 Load'}
                                </button>
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
};

export default ModelManager;