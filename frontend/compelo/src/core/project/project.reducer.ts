import { Action, createReducer, on } from '@ngrx/store';
import { Game, GameStats, Match, Player, PlayerStats } from '@shared/models';
import {
  filterMatches,
  loadGamesSuccess,
  loadGameStatsSuccess,
  loadMatchesSuccess,
  loadPlayersSuccess,
  loadPlayerStatsSuccess
} from './project.actions';

export const FEATURE_KEY = 'project';

export interface State {
  games: Game[];
  gameStats: GameStats;
  players: Player[];
  playerStats: PlayerStats[];
  matches: Match[];
  matchFilter: string;
}

export const initialState: State = {
  games: [],
  gameStats: null,
  players: [],
  playerStats: [],
  matches: [],
  matchFilter: ''
};

const projectReducer = createReducer(
  initialState,
  on(loadGamesSuccess, (state, action) => ({
    ...state,
    games: action.payload || []
  })),
  on(loadGameStatsSuccess, (state, action) => ({
    ...state,
    gameStats: action.payload || null
  })),
  on(loadPlayersSuccess, (state, action) => ({
    ...state,
    players: action.payload || []
  })),
  on(loadPlayerStatsSuccess, (state, action) => ({
    ...state,
    playerStats: action.payload || []
  })),
  on(loadMatchesSuccess, (state, action) => ({
    ...state,
    matches: action.payload || []
  })),
  on(filterMatches, (state, action) => ({
    ...state,
    matchFilter: action.payload.filter || ''
  }))
);

export function reducer(state: State | undefined, action: Action) {
  return projectReducer(state, action);
}
