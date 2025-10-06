import React from 'react';

const Transactions: React.FC = () => {
  return (
    <div>
      <div className="dashboard-header">
        <h1 className="dashboard-title">Transactions 💰</h1>
        <p className="dashboard-subtitle">Track your income and expenses</p>
      </div>
      
      <div className="empty-state">
        <div className="empty-state-icon">💸</div>
        <h3>Transactions Management</h3>
        <p>This page is under development</p>
        <p style={{ fontSize: '14px', marginTop: '8px' }}>
          Soon you'll be able to add and manage transactions here
        </p>
      </div>
    </div>
  );
};

export default Transactions;