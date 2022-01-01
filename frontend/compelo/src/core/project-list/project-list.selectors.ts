import { createFeatureSelector, createSelector } from '@ngrx/store';

import { getSelectedProjectGuid } from '@core/router';

import { FEATURE_KEY, State } from './project-list.reducer';

export const getProjectListState = createFeatureSelector<State>(FEATURE_KEY);

export const getProjects = createSelector(
  getProjectListState,
  (state: State) => state.projects
);

export const getSelectedProject = createSelector(
  getSelectedProjectGuid,
  getProjectListState,
  (guid, state) => {
    return state.projects.find((value) => value.guid === guid);
  }
);
