import { Action, createReducer, on } from '@ngrx/store';
import { Project } from '../../shared/models';
import {
  createProjectSuccess,
  loadProjectsSuccess,
  selectProjectSuccess
} from './project.actions';

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
  on(createProjectSuccess, (state, project) => ({
    ...state,
    selectedProject: project
  })),
  on(selectProjectSuccess, (state, project) => ({
    ...state,
    selectedProject: project
  })),
  on(loadProjectsSuccess, (state, projects) => ({
    ...state,
    projects: projects
  }))
);

export function reducer(state: State | undefined, action: Action) {
  return projectReducer(state, action);
}
