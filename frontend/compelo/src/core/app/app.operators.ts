import { TypedAction } from '@ngrx/store/src/models';
import { Payload } from '@shared/models';
import { filter, map } from 'rxjs/operators';
import { AppActions } from './app.actions';

const __isAppAction = (action: TypedAction<string>) =>
  action.type !== '[App] Loading' && action.type !== '[App] Loaded';

const __escapedActionType = ({ type }: TypedAction<string>) =>
  type?.replace('Success', '')?.replace('Error', '').trim();

export const ofTypeLoading = () =>
  filter(
    (action: Payload<any> & TypedAction<string>) =>
      __isAppAction(action) &&
      !action.type?.toUpperCase()?.startsWith('@NGRX') &&
      !action.type?.toUpperCase()?.includes('SUCCESS') &&
      !action.type?.toUpperCase()?.includes('ERROR')
  );

export const dispatchLoadingAction = () =>
  map(({ type }) => AppActions.loading({ payload: type }));

export const ofTypeLoaded = () =>
  filter(
    (action: Payload<any> & TypedAction<string>) =>
      __isAppAction(action) &&
      !action.type?.toUpperCase()?.startsWith('@NGRX') &&
      (action.type?.toUpperCase()?.endsWith('SUCCESS') ||
        action.type?.toUpperCase()?.includes('ERROR'))
  );

export const dispatchLoadedAction = () =>
  map((action: TypedAction<string>) => {
    return AppActions.loaded({ payload: __escapedActionType(action) });
  });
