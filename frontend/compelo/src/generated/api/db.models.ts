/* Do not change, this code is generated from Golang structs */


export interface Game {
  id: number;
  name: string;
  projectId: number;
}
export interface Match {
  id: number;
  date: string;
  gameId: number;
}
export interface Team {
  id: number;
  matchId: number;
  score: number;
  result: string;
  ratingDelta: number;
}
export interface Appearance {
  id: number;
  matchId: number;
  teamId: number;
  playerId: number;
  ratingDelta: number;
}
export interface MatchResult {
  playerId: number;
  gameId: number;
  date: string;
  matchId: number;
  score: number;
  ratingDelta: number;
  result: string;
}
export interface MatchScoreStats {
  matchId: number;
  gameId: number;
  scoreDiff: number;
  scoreSum: number;
}
export interface Player {
  id: number;
  name: string;
  projectId: number;
}
export interface Project {
  id: number;
  name: string;
}
export interface Rating {
  id: number;
  rating: number;
  gameId: number;
  playerId: number;
}