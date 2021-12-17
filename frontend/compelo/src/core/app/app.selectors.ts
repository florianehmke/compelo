import { createFeatureSelector, createSelector } from '@ngrx/store';
import { appFeatureKey, State } from './app.reducer';

const getAppState = createFeatureSelector<State>(appFeatureKey);

export const getLoading = createSelector(
  getAppState,
  (state: State) => state.loading
);
