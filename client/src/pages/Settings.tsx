import React from 'react';

const Settings: React.FC = () => {
  return (
    <div>
      <div className="dashboard-header">
        <h1 className="dashboard-title">Settings ⚙️</h1>
        <p className="dashboard-subtitle">Configure your application preferences</p>
      </div>
      
      <div className="empty-state">
        <div className="empty-state-icon">⚙️</div>
        <h3>Settings Management</h3>
        <p>This page is under development</p>
        <p style={{ fontSize: '14px', marginTop: '8px' }}>
          Soon you'll be able to customize your settings here
        </p>
      </div>
    </div>
  );
};

export default Settings;