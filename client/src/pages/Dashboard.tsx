import React, { useState, useEffect } from 'react';
import { useAuth } from '../contexts/AuthContext';
import { accountService } from '../services/accountService';
import { transactionService } from '../services/transactionService';
import type { Account, Transaction } from '../types';

const Dashboard: React.FC = () => {
  const { user } = useAuth();
  const [accounts, setAccounts] = useState<Account[]>([]);
  const [transactions, setTransactions] = useState<Transaction[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    try {
      const [accountsData, transactionsData] = await Promise.all([
        accountService.getAll(),
        transactionService.getAll()
      ]);
      setAccounts(accountsData);
      setTransactions(transactionsData);
    } catch (error) {
      console.error('Failed to load data:', error);
    } finally {
      setLoading(false);
    }
  };

  const totalBalance = accounts.reduce((sum, account) => sum + account.balance, 0);
  const totalIncome = transactions
    .filter(t => t.type === 'income')
    .reduce((sum, t) => sum + t.amount, 0);
  const totalExpenses = transactions
    .filter(t => t.type === 'expense')
    .reduce((sum, t) => sum + t.amount, 0);
  
  const recentTransactions = transactions.slice(0, 5);
  const recentAccounts = accounts.slice(0, 3);

  const getBalanceChange = () => {
    if (transactions.length === 0) return 'No data yet';
    return 'Track your progress';
  };

  const getIncomeChange = () => {
    if (totalIncome === 0) return 'Start adding income';
    return 'Growing steadily';
  };

  const getExpenseChange = () => {
    if (totalExpenses === 0) return 'No expenses yet';
    return 'Manage your spending';
  };

  if (loading) {
    return (
      <div className="loading-state">
        <div className="loading-spinner"></div>
        <p>Loading your financial data...</p>
      </div>
    );
  }

  return (
    <div className="dashboard">
      <div className="dashboard-header">
        <h1 className="dashboard-title">Welcome back, {user?.name}! 👋</h1>
        <p className="dashboard-subtitle">
          {accounts.length === 0 ? 'Get started by creating your first account' : 'Here\'s your financial overview'}
        </p>
      </div>

      {/* Stats Grid */}
      <div className="stats-grid">
        <div className="stat-card">
          <div className="stat-header">
            <div className="stat-icon balance">
              ₸
            </div>
          </div>
          <div className="stat-title">Total Balance</div>
          <div className="stat-value">${totalBalance.toFixed(2)}</div>
          <div className="stat-change">{getBalanceChange()}</div>
        </div>

        <div className="stat-card">
          <div className="stat-header">
            <div className="stat-icon income">
              📈
            </div>
          </div>
          <div className="stat-title">Total Income</div>
          <div className="stat-value">${totalIncome.toFixed(2)}</div>
          <div className="stat-change">{getIncomeChange()}</div>
        </div>

        <div className="stat-card">
          <div className="stat-header">
            <div className="stat-icon expense">
              📉
            </div>
          </div>
          <div className="stat-title">Total Expenses</div>
          <div className="stat-value">${totalExpenses.toFixed(2)}</div>
          <div className="stat-change">{getExpenseChange()}</div>
        </div>

        <div className="stat-card">
          <div className="stat-header">
            <div className="stat-icon accounts">
              💳
            </div>
          </div>
          <div className="stat-title">Active Accounts</div>
          <div className="stat-value">{accounts.length}</div>
          <div className="stat-change">
            {accounts.length === 0 ? 'Create first account' : 'All accounts active'}
          </div>
        </div>
      </div>

      {/* Recent Activity */}
      <div className="dashboard-sections">
        {/* Recent Transactions */}
        <div className="section-card">
          <div className="section-header">
            <h3 className="section-title">Recent Transactions</h3>
            <button className="section-action" onClick={() => window.location.href = '/transactions'}>
              View All
            </button>
          </div>
          <div className="section-content">
            {recentTransactions.length > 0 ? (
              recentTransactions.map((transaction) => (
                <div key={transaction.id} className="transaction-item">
                  <div className="transaction-info">
                    <div className={`transaction-icon ${transaction.type}`}>
                      {transaction.type === 'income' ? '⬇️' : '⬆️'}
                    </div>
                    <div className="transaction-details">
                      <h4>{transaction.description || 'No description'}</h4>
                      <p>{new Date(transaction.date).toLocaleDateString()}</p>
                    </div>
                  </div>
                  <div className={`transaction-amount ${transaction.type}`}>
                    {transaction.type === 'income' ? '+' : '-'}${transaction.amount.toFixed(2)}
                  </div>
                </div>
              ))
            ) : (
              <div className="empty-state">
                <div className="empty-state-icon">💸</div>
                <p>No transactions yet</p>
                <p style={{ fontSize: '14px', marginTop: '8px' }}>
                  Start by adding your first transaction
                </p>
                <button 
                  className="btn-modern"
                  style={{ marginTop: '16px' }}
                  onClick={() => window.location.href = '/transactions'}
                >
                  + Add Transaction
                </button>
              </div>
            )}
          </div>
        </div>

        {/* Accounts Overview */}
        <div className="section-card">
          <div className="section-header">
            <h3 className="section-title">Your Accounts</h3>
            <button className="section-action" onClick={() => window.location.href = '/accounts'}>
              View All
            </button>
          </div>
          <div className="section-content">
            {recentAccounts.length > 0 ? (
              recentAccounts.map((account) => (
                <div key={account.id} className="account-item">
                  <div className="account-info">
                    <h4>{account.name}</h4>
                    <p>{account.currency} • Balance: ${account.balance.toFixed(2)}</p>
                  </div>
                  <div className="account-balance">
                    ${account.balance.toFixed(2)}
                  </div>
                </div>
              ))
            ) : (
              <div className="empty-state">
                <div className="empty-state-icon">🏦</div>
                <p>No accounts yet</p>
                <p style={{ fontSize: '14px', marginTop: '8px' }}>
                  Create your first account to get started
                </p>
                <button 
                  className="btn-modern"
                  style={{ marginTop: '16px' }}
                  onClick={() => window.location.href = '/accounts'}
                >
                  + Add Account
                </button>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;