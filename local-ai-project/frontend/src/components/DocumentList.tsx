// frontend/src/components/DocumentList.tsx
import React from 'react';
import type { Document } from '../types/Document';

interface DocumentListProps {
    documents: Document[];
    onRefresh: () => Promise<void>;
}

const DocumentList: React.FC<DocumentListProps> = ({ documents, onRefresh }) => {
    const formatFileSize = (bytes: number): string => {
        const sizes = ['Bytes', 'KB', 'MB', 'GB'];
        if (bytes === 0) return '0 Bytes';
        const i = Math.floor(Math.log(bytes) / Math.log(1024));
        return Math.round(bytes / Math.pow(1024, i) * 100) / 100 + ' ' + sizes[i];
    };

    const formatDate = (dateString: string): string => {
        return new Date(dateString).toLocaleDateString();
    };

    const getFileIcon = (type: string): string => {
        switch (type.toLowerCase()) {
            case 'pdf': return '📄';
            case 'txt': return '📝';
            case 'docx': return '📋';
            case 'md': return '📑';
            default: return '📄';
        }
    };

    return (
        <div className="document-list">
            <div className="section-header">
                <h3>📁 Uploaded Documents ({documents.length})</h3>
                <button onClick={onRefresh} className="refresh-button">
                    🔄 Refresh
                </button>
            </div>

            {documents.length === 0 ? (
                <div className="empty-state">
                    <p>No documents uploaded yet.</p>
                    <p>Upload some documents to get started!</p>
                </div>
            ) : (
                <div className="documents-grid">
                    {documents.map((doc) => (
                        <div key={doc.id} className="document-card">
                            <div className="document-icon">
                                {getFileIcon(doc.type)}
                            </div>
                            <div className="document-info">
                                <h4 className="document-name" title={doc.name}>
                                    {doc.name}
                                </h4>
                                <div className="document-meta">
                                    <span className="file-size">{formatFileSize(doc.size)}</span>
                                    <span className="file-type">{doc.type.toUpperCase()}</span>
                                    <span className="upload-date">{formatDate(doc.uploadDate)}</span>
                                </div>
                                <div className="document-status">
                                    <span className={`status ${doc.status}`}>
                                        {doc.status}
                                    </span>
                                </div>
                            </div>
                        </div>
                    ))}
                </div>
            )}
        </div>
    );
};

export default DocumentList;