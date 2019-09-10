import { createAction, props } from '@ngrx/store';
import {
  CreateProjectPayload,
  ErrorPayload,
  LoginProjectPayload,
  Project
} from '../../shared/models';

export const createProject = createAction(
  '[Project] Create',
  props<CreateProjectPayload>()
);
export const createProjectSuccess = createAction(
  '[Project] Create Success',
  props<Project>()
);
export const createProjectError = createAction(
  '[Project] Create Error',
  props<ErrorPayload>()
);

export const selectProject = createAction(
  '[Project] Select',
  props<LoginProjectPayload>()
);
export const selectProjectSuccess = createAction(
  '[Project] Select Success',
  props<Project>()
);
export const selectProjectError = createAction(
  '[Project] Select Error',
  props<ErrorPayload>()
);

export const loadProjects = createAction('[Project] Load');
export const loadProjectsSuccess = createAction(
  '[Project] Loaded Success',
  props<Project[]>()
);
export const loadProjectsError = createAction(
  '[Project] Load Error',
  props<ErrorPayload>()
);
