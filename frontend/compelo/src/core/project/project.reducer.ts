import { Action, createReducer, on } from '@ngrx/store';
import { Game, Match, Player } from '@shared/models';
import {
  loadGamesSuccess,
  loadMatchesSuccess,
  loadPlayersSuccess
} from './project.actions';

export const FEATURE_KEY = 'project';

export interface State {
  games: Game[];
  players: Player[];
  matches: Match[];
}

export const initialState: State = {
  games: [],
  players: [],
  matches: []
};

const projectReducer = createReducer(
  initialState,
  on(loadGamesSuccess, (state, action) => ({
    ...state,
    games: action.payload
  })),
  on(loadPlayersSuccess, (state, action) => ({
    ...state,
    players: action.payload
  })),
  on(loadMatchesSuccess, (state, action) => ({
    ...state,
    matches: action.payload
  }))
);

export function reducer(state: State | undefined, action: Action) {
  return projectReducer(state, action);
}
