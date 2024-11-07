import { getRouterSelectors, RouterReducerState } from '@ngrx/router-store';
import { createFeatureSelector } from '@ngrx/store';

import { gameIdParam, projectIdParam } from '@shared/route-params';

export const selectRouter =
  createFeatureSelector<RouterReducerState<never>>('router');

const { selectRouteParam } = getRouterSelectors(selectRouter);

export const getSelectedGameId = selectRouteParam(gameIdParam);
export const getSelectedProjectId = selectRouteParam(projectIdParam);
