import { createFeatureSelector, createSelector } from '@ngrx/store';
import { TypedAction } from '@ngrx/store/src/models';
import { appFeatureKey, State } from './app.reducer';

const getAppState = createFeatureSelector<State>(appFeatureKey);

export const getLoading = createSelector(
  getAppState,
  (state: State) => state.loading
);

export const getLoadingBy = ({ type }: TypedAction<string>) => {
  return createSelector(getLoading, (loading) => loading[type]);
};

export const getLoadedBy = ({ type }: TypedAction<string>) => {
  return createSelector(getLoading, (loading) => !loading[type]);
};
