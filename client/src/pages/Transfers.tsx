import React from 'react';

const Transfers: React.FC = () => {
  return (
    <div>
      <div className="dashboard-header">
        <h1 className="dashboard-title">Transfers 🔄</h1>
        <p className="dashboard-subtitle">Transfer money between accounts</p>
      </div>
      
      <div className="empty-state">
        <div className="empty-state-icon">🔄</div>
        <h3>Transfers Management</h3>
        <p>This page is under development</p>
        <p style={{ fontSize: '14px', marginTop: '8px' }}>
          Soon you'll be able to transfer money between your accounts
        </p>
      </div>
    </div>
  );
};

export default Transfers;