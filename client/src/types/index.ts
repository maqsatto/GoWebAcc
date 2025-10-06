export interface User {
  id: number;
  name: string;
  email: string;
  created_at: string;
}

export interface Account {
  id: number;
  user_id: number;
  name: string;
  balance: number;
  currency: string;
  created_at: string;
  transactions?: Transaction[];
}

export interface Category {
  id: number;
  user_id: number;
  name: string;
  type: 'income' | 'expense';
  color: string;
  icon: string;
  created_at: string;
}

export interface Transaction {
  id: number;
  user_id: number;
  account_id: number;
  amount: number;
  description: string;
  category_id: number;
  type: 'income' | 'expense';
  date: string;
  created_at: string;
  category?: Category;
  account?: Account;
}

export interface Transfer {
  id: number;
  user_id: number;
  from_account_id: number;
  to_account_id: number;
  amount: number;
  description: string;
  currency: string;
  exchange_rate: number;
  converted_amount: number;
  date: string;
  created_at: string;
  from_account: Account;
  to_account: Account;
}

export interface Template {
  id: number;
  user_id: number;
  name: string;
  description: string;
  amount: number;
  category_id: number;
  account_id: number;
  created_at: string;
}

export interface AppSettings {
  id: number;
  user_id: number;
  key: string;
  value: string;
}

// API Responses
export interface ApiResponse<T = any> {
  success: boolean;
  message: string;
  data: T;
  error?: string;
}

// Auth
export interface LoginData {
  email: string;
  password: string;
}

export interface RegisterData {
  name: string;
  email: string;
  password: string;
}

export interface AuthResponse {
  token: string;
  user: User;
}