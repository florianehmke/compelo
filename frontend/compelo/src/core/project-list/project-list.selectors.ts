import { createFeatureSelector, createSelector } from '@ngrx/store';
import { FEATURE_KEY, State } from './project-list.reducer';

export const getProjectListState = createFeatureSelector<State>(FEATURE_KEY);

export const getProjects = createSelector(
  getProjectListState,
  (state: State) => state.projects
);
