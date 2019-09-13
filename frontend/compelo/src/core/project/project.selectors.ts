import { createFeatureSelector, createSelector } from '@ngrx/store';
import { FEATURE_KEY, State } from './project.reducer';

export const getProjectViewState = createFeatureSelector<State>(FEATURE_KEY);

export const getSelectedGame = createSelector(
  getProjectViewState,
  (state: State) => state.selectedGame
);

export const getGames = createSelector(
  getProjectViewState,
  (state: State) => state.games
);

export const getPlayers = createSelector(
    getProjectViewState,
    (state: State) => state.players
);
