import { Project } from '@api';
import { Action, createReducer, on } from '@ngrx/store';

import { loadProjectsSuccess } from './project-list.actions';

export const FEATURE_KEY = 'project-list';

export interface State {
  projects: Project[];
}

export const initialState: State = {
  projects: [],
};

const projectListReducer = createReducer(
  initialState,
  on(loadProjectsSuccess, (state, action) => ({
    ...state,
    projects: action.payload,
  }))
);

export function reducer(state: State | undefined, action: Action) {
  return projectListReducer(state, action);
}
