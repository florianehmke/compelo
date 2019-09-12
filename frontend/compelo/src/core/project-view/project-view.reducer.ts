import { Action, createReducer, on } from '@ngrx/store';
import { Game, Player } from '../../shared/models';
import {
  loadGamesSuccess,
  loadPlayersSuccess,
  selectGame
} from './project-view.actions';

export const FEATURE_KEY = 'games';

export interface State {
  games: Game[];
  selectedGame: Game;
  players: Player[];
}

export const initialState: State = {
  games: [],
  selectedGame: null,
  players: []
};

const projectViewReducer = createReducer(
  initialState,
  on(selectGame, (state, action) => ({
    ...state,
    selectedGame: action.payload
  })),
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
  return projectViewReducer(state, action);
}
