export interface CreateProjectPayload {
    name: string;
    password: string;
}

export interface SelectProjectPayload {
    id?: number;
    name: string;
    password: string;
}

export interface TokenPayload {
    token: string;
}
