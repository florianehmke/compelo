export interface Project {
    id?: number;
    name: string;
}

export interface Game {
    id?: number;
    projectId?: number;
    name: string;
}

export interface Player {
    id?: number;
    projectId?: number;
    name: string;
}
