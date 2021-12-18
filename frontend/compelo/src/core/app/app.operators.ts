import { TypedAction } from '@ngrx/store/src/models';
import { filter, map } from 'rxjs/operators';

import { Payload } from '@shared/models';

import { AppActions } from './app.actions';
import {
  DEFAULT_ACTION_TYPES,
  ERROR_ACTION_TYPE,
  SUCCESS_ACTION_TYPE,
} from './app.models';

const __isAppAction = (action: TypedAction<string>) =>
  action.type !== '[App] Loading' && action.type !== '[App] Loaded';

const __escapedActionType = ({ type }: TypedAction<string>) =>
  type?.replace('Success', '')?.replace('Error', '').trim();

const __exclude =
  (excludes: string[]): ((action: TypedAction<string>) => boolean) =>
  ({ type }: TypedAction<string>): boolean =>
    excludes
      .map((exclude) => exclude?.toUpperCase())
      .every((exclude) => !type.toUpperCase().includes(exclude));

const __include =
  (includes: string[]): ((action: TypedAction<string>) => boolean) =>
  ({ type }: TypedAction<string>): boolean =>
    includes
      .map((include) => include?.toUpperCase())
      .some((include) => type.toUpperCase().includes(include));

export const ofTypeLoading = () =>
  filter(
    (action: Payload<any> & TypedAction<string>) =>
      __isAppAction(action) &&
      __exclude([
        ...DEFAULT_ACTION_TYPES,
        SUCCESS_ACTION_TYPE,
        ERROR_ACTION_TYPE,
      ])(action)
  );

export const dispatchLoadingAction = () =>
  map(({ type }) => AppActions.loading({ payload: type }));

export const ofTypeLoaded = () =>
  filter(
    (action: Payload<any> & TypedAction<string>) =>
      __isAppAction(action) &&
      __exclude(DEFAULT_ACTION_TYPES)(action) &&
      __include([SUCCESS_ACTION_TYPE, ERROR_ACTION_TYPE])(action)
  );

export const dispatchLoadedAction = () =>
  map((action: TypedAction<string>) => {
    return AppActions.loaded({ payload: __escapedActionType(action) });
  });
