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

  rating?: number;
  gameCount?: number;
  peakRating?: number;
  lowestRating?: number;
  ratingProgress?: Rating[];
}

export interface Rating {
  rating: number;
  date: string;
}

export interface Match {
  id: number;
  date: string;
  gameId: number;
  teams: Team[];
}

export interface Team {
  score: number;
  result: 'Win' | 'Loss' | 'Draw';
  ratingDelta: number;
  players: Player[];
}
