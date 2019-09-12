export interface Project {
  id?: number;
  name: string;
}

export interface CreateProjectPayload {
  name: string;
  password: string;
}

export interface SelectProjectPayload extends Project {
  password: string;
}

export interface TokenPayload {
  token: string;
}
