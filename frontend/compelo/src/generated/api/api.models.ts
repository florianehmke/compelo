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

export interface Response {
  guid: string;
}
export interface Time {

}
export interface Project {
  id: number;
  createdDate: Time;
  updatedDate: Time;
  guid: string;
  name: string;
  passwordHash: number[];
}
export interface Player {
  id: number;
  createdDate: Time;
  updatedDate: Time;
  guid: string;
  projectGuid: string;
  name: string;
}
export interface Game {
  id: number;
  createdDate: Time;
  updatedDate: Time;
  guid: string;
  projectGuid: string;
  name: string;
}
export interface MatchTeam {
  players: Player[];
  score: number;
  result: string;
  ratingDelta: number;
}
export interface Match {
  id: number;
  createdDate: Time;
  updatedDate: Time;
  guid: string;
  gameGuid: string;
  projectGuid: string;
  date: string;
  teams: MatchTeam[];
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
  id: number;
  createdDate: Time;
  updatedDate: Time;
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