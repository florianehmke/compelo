import { createFeatureSelector, createSelector } from '@ngrx/store';
import { getSelectedGameId } from '@core/router';

import { FEATURE_KEY, State } from './project.reducer';

export const getProjectState = createFeatureSelector<State>(FEATURE_KEY);

export const getSelectedGame = createSelector(
  getSelectedGameId,
  getProjectState,
  (id, state) => {
    return state.games.find(value => value.id === parseInt(id, 10));
  }
);

export const getGames = createSelector(
  getProjectState,
  (state: State) => state.games
);

export const getPlayers = createSelector(
  getProjectState,
  (state: State) => state.players
);

export const getPlayerStats = createSelector(
  getProjectState,
  (state: State) => state.playerStats
);

export const getMatches = createSelector(
  getProjectState,
  (state: State) =>
    state.matches.filter(match => {
      const searchTerm = state.matchFilter.toLowerCase();
      return match.teams.some(team => {
        return team.players.some(player => {
          return player.name.toLowerCase().includes(searchTerm);
        });
      });
    })
);
