import { api } from './api';
import type { Transaction, ApiResponse } from '../types';

export const transactionService = {
  async getAll(): Promise<Transaction[]> {
    const response = await api.get<ApiResponse<Transaction[]>>('/transactions');
    return response.data.data;
  },

  async getById(id: number): Promise<Transaction> {
    const response = await api.get<ApiResponse<Transaction>>(`/transactions/${id}`);
    return response.data.data;
  },

  async create(data: Omit<Transaction, 'id' | 'user_id' | 'created_at'>): Promise<Transaction> {
    const response = await api.post<ApiResponse<Transaction>>('/transactions', data);
    return response.data.data;
  },

  async update(id: number, data: Partial<Transaction>): Promise<Transaction> {
    const response = await api.put<ApiResponse<Transaction>>(`/transactions/${id}`, data);
    return response.data.data;
  },

  async delete(id: number): Promise<void> {
    await api.delete(`/transactions/${id}`);
  },
};