/* Do not change, this code is generated from Golang structs */


export interface PlayerData {
  id: number;
  name: string;
  projectId: number;
}
export interface TeamData {
  id: number;
  matchId: number;
  score: number;
  result: string;
  ratingDelta: number;
  players: PlayerData[];
}
export interface MatchData {
  id: number;
  date: string;
  gameId: number;
  teams: TeamData[];
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
  name: string;
  projectId: number;
  current: Stats;
  history: {[key: string]: Stats};
}
export interface GameStats {
  maxScoreSum: MatchData[];
  maxScoreDiff: MatchData[];
}