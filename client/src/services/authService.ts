import { api } from './api';
import type { LoginData, RegisterData, AuthResponse, ApiResponse, User } from '../types';

export const authService = {
  async login(data: LoginData): Promise<AuthResponse> {
    const response = await api.post<ApiResponse<AuthResponse>>('/users/login', data);
    return response.data.data;
  },

  async register(data: RegisterData): Promise<User> {
    const response = await api.post<ApiResponse<User>>('/users/register', data);
    return response.data.data;
  },

  async getProfile(): Promise<User> {
    const response = await api.get<ApiResponse<User>>('/users/profile');
    return response.data.data;
  },
};