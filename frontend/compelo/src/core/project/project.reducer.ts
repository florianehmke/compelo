import { Action, createReducer, on } from '@ngrx/store';
import { Project } from '../../shared/models';
import {loadProjectsSuccess, selectProject, selectProjectSuccess} from './project.actions';

export const FEATURE_KEY = 'project';

export interface State {
  projects: Project[];
  selectedProject: Project;
}

export const initialState: State = {
  projects: [],
  selectedProject: null
};

const projectReducer = createReducer(
  initialState,
  on(selectProjectSuccess, (state, action) => ({
    ...state,
    selectedProject: action.payload
  })),
  on(loadProjectsSuccess, (state, action) => ({
    ...state,
    projects: action.payload
  }))
);

export function reducer(state: State | undefined, action: Action) {
  return projectReducer(state, action);
}
