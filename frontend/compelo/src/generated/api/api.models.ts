/* Do not change, this code is generated from Golang structs */


export interface AuthRequest {
  projectId: number;
  password: string;
}
export interface AuthResponse {
  token: string;
}
export interface CreateProjectRequest {
  name: string;
  password: string;
}
export interface CreateGameRequest {
  name: string;
}
export interface CreatePlayerRequest {
  name: string;
}
export interface CreateMatchRequestTeam {
  playerIds: number[];
  score: number;
}
export interface CreateMatchRequest {
  teams: CreateMatchRequestTeam[];
}
