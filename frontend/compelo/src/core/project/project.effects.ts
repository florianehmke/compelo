import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { Store } from '@ngrx/store';
import { catchError, map, switchMap, withLatestFrom } from 'rxjs/operators';
import { of } from 'rxjs';

import { ProjectService } from './project.service';
import {
  createGame,
  createGameError,
  createMatch,
  createMatchError,
  createMatchSuccess,
  createPlayer,
  createPlayerError,
  loadGames,
  loadGamesError,
  loadGamesSuccess,
  loadPlayers,
  loadPlayersError,
  loadPlayersSuccess
} from './project.actions';
import { State } from './project.reducer';
import { getSelectedGame } from './project.selectors';

@Injectable()
export class ProjectEffects {
  loadGames$ = createEffect(() =>
    this.actions$.pipe(
      ofType(loadGames),
      switchMap(() =>
        this.service.getGames().pipe(
          map(games => loadGamesSuccess({ payload: games })),
          catchError(err => of(loadGamesError(err)))
        )
      )
    )
  );

  createGame$ = createEffect(() =>
    this.actions$.pipe(
      ofType(createGame),
      switchMap(action =>
        this.service.createGame(action.payload).pipe(
          map(() => loadGames()),
          catchError(err => of(createGameError(err)))
        )
      )
    )
  );

  loadPlayers$ = createEffect(() =>
    this.actions$.pipe(
      ofType(loadPlayers),
      switchMap(() =>
        this.service.getPlayers().pipe(
          map(players => loadPlayersSuccess({ payload: players })),
          catchError(err => of(loadPlayersError(err)))
        )
      )
    )
  );

  createPlayer$ = createEffect(() =>
    this.actions$.pipe(
      ofType(createPlayer),
      switchMap(action =>
        this.service.createPlayer(action.payload).pipe(
          map(() => loadPlayers()),
          catchError(err => of(createPlayerError(err)))
        )
      )
    )
  );

  createMatch$ = createEffect(() =>
    this.actions$.pipe(
      ofType(createMatch),
      withLatestFrom(this.store.select(getSelectedGame)),
      switchMap(([action, game]) =>
        this.service.createMatch(action.payload, game.id).pipe(
          map(response => createMatchSuccess({ payload: response })),
          catchError(err => of(createMatchError(err)))
        )
      )
    )
  );

  constructor(
    private actions$: Actions,
    private service: ProjectService,
    private store: Store<State>
  ) {}
}
