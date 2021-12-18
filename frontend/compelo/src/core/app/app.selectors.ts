import { createFeatureSelector, createSelector } from '@ngrx/store';
import { ActionCreator } from '@ngrx/store/src/models';

import { appFeatureKey, State } from './app.reducer';

const getAppState = createFeatureSelector<State>(appFeatureKey);

export const getLoading = createSelector(
  getAppState,
  (state: State) => state.loading
);

export const getLoadingByActionTypeOf = (
  actionCreator: ActionCreator<string, any>
) => {
  const { type: actionType } = actionCreator();
  return createSelector(getLoading, (loading) => loading[actionType]);
};

export const getLoadedByActionTypeOf = (
  actionCreator: ActionCreator<string, any>
) => {
  const { type: actionType } = actionCreator();
  return createSelector(getLoading, (loading) => !loading[actionType]);
};
