import React from 'react';

const Categories: React.FC = () => {
  return (
    <div>
      <div className="dashboard-header">
        <h1 className="dashboard-title">Categories 🏷️</h1>
        <p className="dashboard-subtitle">Organize your transactions with categories</p>
      </div>
      
      <div className="empty-state">
        <div className="empty-state-icon">🏷️</div>
        <h3>Categories Management</h3>
        <p>This page is under development</p>
        <p style={{ fontSize: '14px', marginTop: '8px' }}>
          Soon you'll be able to create and manage categories here
        </p>
      </div>
    </div>
  );
};

export default Categories;