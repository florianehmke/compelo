import { createAction, props } from '@ngrx/store';
import { Payload } from '@shared/models';

const loading = createAction('[App] Loading', props<Payload<string>>());

const loaded = createAction('[App] Loaded', props<Payload<string>>());

export const AppActions = {
  loading,
  loaded,
};
