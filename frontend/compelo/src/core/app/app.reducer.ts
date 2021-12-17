import { createReducer, on } from '@ngrx/store';
import { AppActions } from './app.actions';
import { LoadingActions } from './app.models';

export const appFeatureKey = 'app';

export interface State {
  loading: LoadingActions;
}

const initialState: State = {
  loading: {},
};

export const reducer = createReducer(
  initialState,
  on(AppActions.loading, (state, { payload }) => ({
    loading: {
      ...state.loading,
      [payload]: true,
    },
  })),
  on(AppActions.loaded, (state, { payload }) => ({
    loading: {
      ...state.loading,
      [payload]: false,
    },
  }))
);
