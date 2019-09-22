export interface CreateProjectPayload {
    name: string;
    password: string;
}

export interface SelectProjectPayload {
    projectId: number;
    projectName: string;
    password: string;
}
