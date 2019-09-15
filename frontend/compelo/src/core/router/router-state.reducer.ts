import * as fromRouter from '@ngrx/router-store';

export const FEATURE_KEY = 'router';

export interface State {
  router: fromRouter.RouterReducerState<any>;
}
