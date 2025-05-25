// frontend/src/components/DocumentUploader.tsx
import React, { useState, useRef } from 'react';

interface DocumentUploaderProps {
    onUpload: (file: File) => void;
}

const DocumentUploader: React.FC<DocumentUploaderProps> = ({ onUpload }) => {
    const [dragActive, setDragActive] = useState(false);
    const [uploading, setUploading] = useState(false);
    const inputRef = useRef<HTMLInputElement>(null);

    const handleDrag = (e: React.DragEvent) => {
        e.preventDefault();
        e.stopPropagation();
        if (e.type === "dragenter" || e.type === "dragover") {
            setDragActive(true);
        } else if (e.type === "dragleave") {
            setDragActive(false);
        }
    };

    const handleDrop = (e: React.DragEvent) => {
        e.preventDefault();
        e.stopPropagation();
        setDragActive(false);

        if (e.dataTransfer.files && e.dataTransfer.files[0]) {
            handleFile(e.dataTransfer.files[0]);
        }
    };

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        e.preventDefault();
        if (e.target.files && e.target.files[0]) {
            handleFile(e.target.files[0]);
        }
    };

    const handleFile = async (file: File) => {
        const allowedTypes = ['.pdf', '.txt', '.docx', '.md'];
        const fileExt = '.' + file.name.split('.').pop()?.toLowerCase();

        if (!allowedTypes.includes(fileExt)) {
            alert('Unsupported file type. Please upload PDF, TXT, DOCX, or MD files.');
            return;
        }

        if (file.size > 10 * 1024 * 1024) { // 10MB limit
            alert('File too large. Maximum size is 10MB.');
            return;
        }

        setUploading(true);
        try {
            await onUpload(file);
        } catch (error) {
            console.error('Upload failed:', error);
        } finally {
            setUploading(false);
            if (inputRef.current) {
                inputRef.current.value = '';
            }
        }
    };

    const onButtonClick = () => {
        inputRef.current?.click();
    };

    return (
        <div className="document-uploader">
            <div
                className={`upload-zone ${dragActive ? 'drag-active' : ''} ${uploading ? 'uploading' : ''}`}
                onDragEnter={handleDrag}
                onDragLeave={handleDrag}
                onDragOver={handleDrag}
                onDrop={handleDrop}
                onClick={onButtonClick}
            >
                <input
                    ref={inputRef}
                    type="file"
                    accept=".pdf,.txt,.docx,.md"
                    onChange={handleChange}
                    style={{ display: 'none' }}
                />

                {uploading ? (
                    <div className="upload-status">
                        <div className="spinner">⏳</div>
                        <p>Uploading document...</p>
                    </div>
                ) : (
                    <div className="upload-content">
                        <div className="upload-icon">📎</div>
                        <h3>Upload Document</h3>
                        <p>Drag and drop a file here, or click to select</p>
                        <p className="file-types">
                            Supported: PDF, TXT, DOCX, MD (max 10MB)
                        </p>
                    </div>
                )}
            </div>
        </div>
    );
};

export default DocumentUploader;

