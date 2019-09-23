export interface CreateMatchPayload {
  teams: {
    playerIds: number[];
    score: number;
  }[];
}

export interface LoadMatchesPayload {
  gameId: number;
}

export interface LoadPlayersWithStatusPayload {
  gameId: number;
}

export interface FilterMatchesPayload {
  filter: string;
}
