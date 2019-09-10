export interface Project {
  id: number;
  name: string;
}

export interface CreateProjectPayload {
  name: string;
  password: string;
}

export interface LoginProjectPayload {
  name: string;
  password: string;
}
