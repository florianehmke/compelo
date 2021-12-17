import { Injectable } from '@angular/core';
import { Actions, createEffect } from '@ngrx/effects';
import {
  dispatchLoadedAction,
  dispatchLoadingAction,
  ofTypeLoaded,
  ofTypeLoading,
} from './app.operators';

@Injectable()
export class AppEffects {
  loading$ = createEffect(() =>
    this._actions$.pipe(ofTypeLoading(), dispatchLoadingAction())
  );

  loaded$ = createEffect(() =>
    this._actions$.pipe(ofTypeLoaded(), dispatchLoadedAction())
  );

  constructor(private _actions$: Actions) {}
}
