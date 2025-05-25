import React from 'react';
import type { WikiResult } from '../types/WikiResult';

interface WikiResultsProps {
  results: WikiResult[];
  isLoading?: boolean;
}

const WikiResults: React.FC<WikiResultsProps> = ({ results, isLoading = false }) => {
  if (isLoading) {
    return (
      <div className="wiki-results">
        <h3>🌐 Wikipedia Results</h3>
        <div className="loading">Searching Wikipedia...</div>
      </div>
    );
  }

  if (!results || results.length === 0) {
    return null;
  }

  return (
    <div className="wiki-results">
      <h3>🌐 Wikipedia Results ({results.length})</h3>
      <div className="wiki-list">
        {results.map((result, index) => (
          <div key={result.pageId || index} className="wiki-item">
            <div className="wiki-content">
              <h4 className="wiki-title">
                <a href={result.url} target="_blank" rel="noopener noreferrer">
                  {result.title}
                </a>
              </h4>
              {result.description && (
                <p className="wiki-description">{result.description}</p>
              )}
              {result.extract && (
                <p className="wiki-extract">{result.extract}</p>
              )}
              {result.relevanceScore && (
                <span className="relevance-score">
                  Relevance: {(result.relevanceScore * 100).toFixed(1)}%
                </span>
              )}
            </div>
            {result.thumbnail && (
              <div className="wiki-thumbnail">
                <img 
                  src={result.thumbnail} 
                  alt={result.title}
                  onError={(e) => {
                    e.currentTarget.style.display = 'none';
                  }}
                />
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  );
};

export default WikiResults;