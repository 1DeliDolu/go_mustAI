// frontend/src/components/ChatInterface.tsx
import React, { useState } from 'react';
import type { QueryResponse } from '../types/chatMessage';

interface ChatInterfaceProps {
    onQuery: (query: string, includeWiki: boolean) => void;
    loading: boolean;
    response: QueryResponse | null;
}

const ChatInterface: React.FC<ChatInterfaceProps> = ({ onQuery, loading, response }) => {
    const [query, setQuery] = useState('');
    const [includeWiki, setIncludeWiki] = useState(true);

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        if (query.trim()) {
            onQuery(query, includeWiki);
        }
    };

    return (
        <div className="chat-interface">
            <div className="chat-history">
                {response && (
                    <div className="message-bubble ai-message">
                        <div className="message-content">
                            {response.response}
                        </div>
                        <div className="message-sources">
                            <small>
                                Sources: {response.sources.documents.length} documents, {response.sources.wiki.length} wiki results
                            </small>
                        </div>
                    </div>
                )}
            </div>

            <form onSubmit={handleSubmit} className="query-form">
                <div className="input-group">
                    <input
                        type="text"
                        value={query}
                        onChange={(e) => setQuery(e.target.value)}
                        placeholder="Ask a question..."
                        className="query-input"
                        disabled={loading}
                    />
                    <button
                        type="submit"
                        disabled={loading || !query.trim()}
                        className="query-button"
                    >
                        {loading ? '⏳' : '🚀'}
                    </button>
                </div>

                <div className="query-options">
                    <label className="checkbox-label">
                        <input
                            type="checkbox"
                            checked={includeWiki}
                            onChange={(e) => setIncludeWiki(e.target.checked)}
                        />
                        Include Wikipedia search
                    </label>
                </div>
            </form>
        </div>
    );
};

export default ChatInterface;

