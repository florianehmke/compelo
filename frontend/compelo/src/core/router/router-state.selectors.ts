import * as fromRouter from '@ngrx/router-store';
import { createFeatureSelector } from '@ngrx/store';

import { gameIdParam, projectIdParam } from '@shared/route-params';

import { FEATURE_KEY, State } from './router-state.reducer';

export const selectRouter = createFeatureSelector<
  State,
  fromRouter.RouterReducerState<any>
>(FEATURE_KEY);

const { selectRouteParam } = fromRouter.getSelectors(selectRouter);

export const getSelectedGameId = selectRouteParam(gameIdParam);
export const getSelectedProjectId = selectRouteParam(projectIdParam);
