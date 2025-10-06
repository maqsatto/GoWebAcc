import { api } from './api';
import type { Account, ApiResponse } from '../types';

export const accountService = {
  async getAll(): Promise<Account[]> {
    const response = await api.get<ApiResponse<Account[]>>('/accounts');
    return response.data.data;
  },

  async getById(id: number): Promise<Account> {
    const response = await api.get<ApiResponse<Account>>(`/accounts/${id}`);
    return response.data.data;
  },

  async create(data: Omit<Account, 'id' | 'user_id' | 'created_at' | 'transactions'>): Promise<Account> {
    const response = await api.post<ApiResponse<Account>>('/accounts', data);
    return response.data.data;
  },

  async update(id: number, data: Partial<Account>): Promise<Account> {
    const response = await api.put<ApiResponse<Account>>(`/accounts/${id}`, data);
    return response.data.data;
  },

  async delete(id: number): Promise<void> {
    await api.delete(`/accounts/${id}`);
  },
};