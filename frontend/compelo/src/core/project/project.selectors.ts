import { createFeatureSelector, createSelector } from '@ngrx/store';
import { FEATURE_KEY, State } from './project.reducer';

export const getProjectState = createFeatureSelector<State>(FEATURE_KEY);

export const getSelectedProject = createSelector(
  getProjectState,
  (state: State) => state.selectedProject
);

export const getProjects = createSelector(
  getProjectState,
  (state: State) => state.projects
);
