export interface CreateMatchPayload {
  teams: {
    playerIds: number[];
    score: number;
  }[];
}

export interface LoadMatchesPayload {
  gameId: number;
}

export interface LoadPlayerStatsPayload {
  gameId: number;
}

export interface LoadGameStatsPayload {
  gameId: number;
}

export interface FilterMatchesPayload {
  filter: string;
}
