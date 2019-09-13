import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { Router } from '@angular/router';
import { catchError, map, switchMap } from 'rxjs/operators';
import { of } from 'rxjs';

import { ProjectService } from './project.service';
import {
  createGame,
  createGameError,
  createPlayer,
  createPlayerError,
  loadGames,
  loadGamesError,
  loadGamesSuccess,
  loadPlayers,
  loadPlayersError,
  loadPlayersSuccess
} from './project.actions';

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

  constructor(
    private actions$: Actions,
    private service: ProjectService,
    private router: Router
  ) {}
}
