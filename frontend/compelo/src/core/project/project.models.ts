export interface CreateMatchPayload {
  teams: Team[];
}

export interface Team {
  playerIds: number[];
  score: number;
}

export interface LoadMatchesPayload {
  gameId: number;
}

export interface LoadPlayersWithStatusPayload {
  gameId: number;
}
