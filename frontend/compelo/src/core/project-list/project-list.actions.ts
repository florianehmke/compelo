import { createAction, props } from '@ngrx/store';
import { ErrorPayload, Payload } from '@shared/models';
import { CreateProjectRequest, Project } from '@api';

import { SelectProjectPayload } from './project-list.models';

export const loadProjects = createAction('[Projects] Load');
export const loadProjectsSuccess = createAction(
  '[Projects] Load Success',
  props<Payload<Project[]>>()
);
export const loadProjectsError = createAction(
  '[Projects] Load Error',
  props<ErrorPayload>()
);

export const createProject = createAction(
  '[Projects] Create',
  props<Payload<CreateProjectRequest>>()
);
export const createProjectSuccess = createAction(
  '[Projects] Create Success',
  props<Payload<Project>>()
);
export const createProjectError = createAction(
  '[Projects] Create Error',
  props<ErrorPayload>()
);

export const selectProject = createAction(
  '[Projects] Select',
  props<Payload<SelectProjectPayload>>()
);
export const selectProjectSuccess = createAction(
  '[Projects] Select Success',
  props<Payload<Project>>()
);
export const selectProjectError = createAction(
  '[Projects] Select Error',
  props<ErrorPayload>()
);
