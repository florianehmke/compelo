/* Do not change, this code is generated from Golang structs */


export interface AuthRequest {
  projectGuid: string;
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
  playerGuids: string[];
  score: number;
}
export interface CreateMatchRequest {
  teams: CreateMatchRequestTeam[];
}

export interface Project {
  guid: string;
  name: string;
  passwordHash: number[];
}
export interface Player {
  guid: string;
  projectGuid: string;
  name: string;
}
export interface Game {
  guid: string;
  projectGuid: string;
  name: string;
}
export interface Team {
  players: Player[];
  score: number;
  result: string;
  ratingDelta: number;
}
export interface Match {
  guid: string;
  gameGuid: string;
  projectGuid: string;
  date: string;
  teams: Team[];
}
export interface Stats {
  rating: number;
  peakRating: number;
  lowestRating: number;
  gameCount: number;
  winCount: number;
  drawCount: number;
  lossCount: number;
}
export interface PlayerStats {
  guid: string;
  projectGuid: string;
  name: string;
  current: Stats;
  history: {[key: string]: Stats};
}
export interface GameStats {
  maxScoreSum: Match[];
  maxScoreDiff: Match[];
}