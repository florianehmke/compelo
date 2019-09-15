import * as fromRouter from '@ngrx/router-store';
import { createFeatureSelector } from '@ngrx/store';
import { FEATURE_KEY, State } from './router-state.reducer';

export const selectRouter = createFeatureSelector<
  State,
  fromRouter.RouterReducerState<any>
>(FEATURE_KEY);

const { selectRouteParam } = fromRouter.getSelectors(selectRouter);

export const getSelectedGameId = selectRouteParam('gameId');
