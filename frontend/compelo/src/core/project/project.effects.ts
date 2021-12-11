import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { Store } from '@ngrx/store';
import {
  catchError,
  map,
  switchMap,
  tap,
  withLatestFrom,
} from 'rxjs/operators';
import { of } from 'rxjs';
import { ToastService } from '@shared/toast';

import { ProjectService } from './project.service';
import {
  createGame,
  createGameError,
  createGameSuccess,
  createMatch,
  createMatchError,
  createMatchSuccess,
  createPlayer,
  createPlayerError,
  createPlayerSuccess,
  loadGames,
  loadGamesError,
  loadGamesSuccess,
  loadGameStats,
  loadGameStatsError,
  loadGameStatsSuccess,
  loadMatches,
  loadMatchesError,
  loadMatchesSuccess,
  loadPlayers,
  loadPlayersError,
  loadPlayersSuccess,
  loadPlayerStats,
  loadPlayerStatsError,
  loadPlayerStatsSuccess,
} from './project.actions';
import { State } from './project.reducer';
import { getSelectedGame } from './project.selectors';

@Injectable()
export class ProjectEffects {
  loadGames$ = createEffect(() =>
    this.actions$.pipe(
      ofType(loadGames),
      switchMap((action) =>
        this.service.getGames().pipe(
          map((games) => loadGamesSuccess({ payload: games })),
          catchError((err) => of(loadGamesError(err)))
        )
      )
    )
  );

  createGame$ = createEffect(() =>
    this.actions$.pipe(
      ofType(createGame),
      switchMap((action) =>
        this.service.createGame(action.payload).pipe(
          switchMap((response) => [
            createGameSuccess({ payload: response }),
            loadGames(),
          ]),
          catchError((err) => of(createGameError(err)))
        )
      )
    )
  );

  loadGameStats$ = createEffect(() =>
    this.actions$.pipe(
      ofType(loadGameStats),
      switchMap((action) =>
        this.service.getGameStats(action.payload.gameId).pipe(
          map((gameStats) => loadGameStatsSuccess({ payload: gameStats })),
          catchError((err) => of(loadGameStatsError(err)))
        )
      )
    )
  );

  loadPlayers$ = createEffect(() =>
    this.actions$.pipe(
      ofType(loadPlayers),
      switchMap(() =>
        this.service.getPlayers().pipe(
          map((players) => loadPlayersSuccess({ payload: players })),
          catchError((err) => of(loadPlayersError(err)))
        )
      )
    )
  );

  loadPlayerStats$ = createEffect(() =>
    this.actions$.pipe(
      ofType(loadPlayerStats),
      switchMap((action) =>
        this.service.getPlayerStats(action.payload.gameId).pipe(
          map((players) => loadPlayerStatsSuccess({ payload: players })),
          catchError((err) => of(loadPlayerStatsError(err)))
        )
      )
    )
  );

  createPlayer$ = createEffect(() =>
    this.actions$.pipe(
      ofType(createPlayer),
      switchMap((action) =>
        this.service.createPlayer(action.payload).pipe(
          switchMap((response) => [
            createPlayerSuccess({ payload: response }),
            loadPlayers(),
          ]),
          catchError((err) => of(createPlayerError(err)))
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
          switchMap((response) => [
            createMatchSuccess({ payload: response }),
            loadMatches({ payload: { gameId: game.id } }),
            loadPlayerStats({ payload: { gameId: game.id } }),
            loadGameStats({ payload: { gameId: game.id } }),
          ]),
          catchError((err) => of(createMatchError(err)))
        )
      )
    )
  );

  loadMatches$ = createEffect(() =>
    this.actions$.pipe(
      ofType(loadMatches),
      switchMap((action) =>
        this.service.getMatches(action.payload.gameId).pipe(
          map((matches) => loadMatchesSuccess({ payload: matches })),
          catchError((err) => of(loadMatchesError(err)))
        )
      )
    )
  );

  notifications$ = createEffect(
    () =>
      this.actions$.pipe(
        ofType(createMatchSuccess, createPlayerSuccess, createGameSuccess),
        tap((action) => this.toastService.success('Created!'))
      ),
    { dispatch: false }
  );

  constructor(
    private actions$: Actions,
    private service: ProjectService,
    private store: Store<State>,
    private toastService: ToastService
  ) {}
}
