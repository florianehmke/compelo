export interface CreateMatchPayload {
  teams: Team[];
}

export interface Team {
  playerIds: number[];
  score: number;
  winner: boolean;
}

export interface LoadGamesPayload {
  loadMatches: boolean;
}
