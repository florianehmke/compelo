import { createFeatureSelector, createSelector } from '@ngrx/store';
import { getSelectedProjectId } from '@core/router';

import { FEATURE_KEY, State } from './project-list.reducer';

export const getProjectListState = createFeatureSelector<State>(FEATURE_KEY);

export const getProjects = createSelector(
  getProjectListState,
  (state: State) => state.projects
);

export const getSelectedProject = createSelector(
  getSelectedProjectId,
  getProjectListState,
  (id, state) => {
    return state.projects.find(value => value.id === parseInt(id, 10));
  }
);
