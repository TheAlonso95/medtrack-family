import Cookies from 'js-cookie';

// Types
export interface User {
  id: string;
  email: string;
  name?: string;
}

export interface LoginCredentials {
  email: string;
  password: string;
}

export interface SignupCredentials {
  email: string;
  password: string;
  name?: string;
}

export interface AuthResponse {
  token: string;
  user: User;
}

// API URL
const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

// Token management
export const getToken = (): string | undefined => {
  return Cookies.get('auth_token');
};

export const setToken = (token: string): void => {
  Cookies.set('auth_token', token, { 
    expires: 7, // 7 days
    secure: process.env.NODE_ENV === 'production',
    sameSite: 'strict'
  });
};

export const removeToken = (): void => {
  Cookies.remove('auth_token');
};

// Auth API calls
export const login = async (credentials: LoginCredentials): Promise<AuthResponse> => {
  try {
    const response = await fetch(`${API_URL}/api/auth/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(credentials),
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.message || 'Failed to login');
    }

    const data = await response.json();
    setToken(data.token);
    return data;
  } catch (error) {
    if (error instanceof Error) {
      throw new Error(error.message);
    }
    throw new Error('An unknown error occurred during login');
  }
};

export const signup = async (credentials: SignupCredentials): Promise<AuthResponse> => {
  try {
    const response = await fetch(`${API_URL}/api/auth/signup`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(credentials),
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.message || 'Failed to signup');
    }

    const data = await response.json();
    setToken(data.token);
    return data;
  } catch (error) {
    if (error instanceof Error) {
      throw new Error(error.message);
    }
    throw new Error('An unknown error occurred during signup');
  }
};

export const logout = (): void => {
  removeToken();
  // Optionally, you can add a call to the backend to invalidate the token
};

export const getCurrentUser = async (): Promise<User | null> => {
  const token = getToken();
  
  if (!token) {
    return null;
  }
  
  try {
    const response = await fetch(`${API_URL}/api/auth/me`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      removeToken();
      return null;
    }

    return await response.json();
  } catch (error) {
    removeToken();
    return null;
  }
};