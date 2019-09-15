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

export interface Match {
    id: number;
    date: string;
    gameId: number;
    teams: Team[]
}

export interface Team {
    score: number;
    winner: boolean;
    players: Player[];
}
