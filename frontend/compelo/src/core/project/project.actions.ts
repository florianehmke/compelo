import { createAction, props } from '@ngrx/store';
import {
  CreateProjectPayload,
  ErrorPayload,
  SelectProjectPayload,
  Project,
  Payload
} from '../../shared/models';

export const createProject = createAction(
  '[Project] Create',
  props<Payload<CreateProjectPayload>>()
);
export const createProjectSuccess = createAction(
  '[Project] Create Success',
  props<Payload<Project>>()
);
export const createProjectError = createAction(
  '[Project] Create Error',
  props<ErrorPayload>()
);

export const selectProject = createAction(
  '[Project] Select',
  props<Payload<SelectProjectPayload>>()
);
export const selectProjectSuccess = createAction(
  '[Project] Select Success',
  props<Payload<Project>>()
);
export const selectProjectError = createAction(
  '[Project] Select Error',
  props<ErrorPayload>()
);

export const loadProjects = createAction('[Project] Load');
export const loadProjectsSuccess = createAction(
  '[Project] Loaded Success',
  props<Payload<Project[]>>()
);
export const loadProjectsError = createAction(
  '[Project] Load Error',
  props<ErrorPayload>()
);
