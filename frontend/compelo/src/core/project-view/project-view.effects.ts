import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { ProjectViewService } from './project-view.service';
import { Router } from '@angular/router';
import { catchError, map, switchMap } from 'rxjs/operators';
import { of } from 'rxjs';
import {
  loadGames,
  loadGamesError,
  loadGamesSuccess,
  loadPlayers,
  loadPlayersError,
  loadPlayersSuccess
} from './project-view.actions';

@Injectable()
export class ProjectViewEffects {
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

  constructor(
    private actions$: Actions,
    private service: ProjectViewService,
    private router: Router
  ) {}
}
