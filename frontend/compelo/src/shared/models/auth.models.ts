export interface AuthRequest {
  projectName: string;
  password: string;
}

export interface TokenPayload {
  token: string;
  expire: string;
}
