import { Action, createReducer, on } from '@ngrx/store';
import { Game, Player } from '@shared/models';
import { loadGamesSuccess, loadPlayersSuccess } from './project.actions';

export const FEATURE_KEY = 'project';

export interface State {
  games: Game[];
  players: Player[];
}

export const initialState: State = {
  games: [],
  players: []
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
  }))
);

export function reducer(state: State | undefined, action: Action) {
  return projectReducer(state, action);
}
