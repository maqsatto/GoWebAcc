import React from 'react';

const Accounts: React.FC = () => {
  return (
    <div>
      <div className="dashboard-header">
        <h1 className="dashboard-title">Accounts 💳</h1>
        <p className="dashboard-subtitle">Manage your financial accounts</p>
      </div>
      
      <div className="empty-state">
        <div className="empty-state-icon">🏦</div>
        <h3>Accounts Management</h3>
        <p>This page is under development</p>
        <p style={{ fontSize: '14px', marginTop: '8px' }}>
          Soon you'll be able to create and manage your accounts here
        </p>
      </div>
    </div>
  );
};

export default Accounts;