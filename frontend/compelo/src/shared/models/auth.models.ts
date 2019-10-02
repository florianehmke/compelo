export interface AuthRequest {
  projectName: string;
  password: string;
}

export interface AuthResponse {
  token: string;
}
