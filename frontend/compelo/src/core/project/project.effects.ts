import { Injectable } from '@angular/core';
import { catchError, map, switchMap, tap } from 'rxjs/operators';
import { of } from 'rxjs';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { ProjectService } from './project.service';
import {
  loadProjects,
  loadProjectsError,
  loadProjectsSuccess,
  selectProject,
  selectProjectSuccess
} from './project.actions';
import { Router } from '@angular/router';

@Injectable()
export class ProjectEffects {
  loadProjects$ = createEffect(() =>
    this.actions$.pipe(
      ofType(loadProjects),
      switchMap(() =>
        this.service.getProjects().pipe(
          map(projects => loadProjectsSuccess({ payload: projects })),
          catchError(err => of(loadProjectsError(err)))
        )
      )
    )
  );

  selectProject$ = createEffect(() =>
    this.actions$.pipe(
      ofType(selectProject),
      switchMap(action =>
        this.service.selectProject(action.payload).pipe(
          tap(r => localStorage.setItem('compelo-token', r.token)),
          tap(() => delete action.payload.password),
          map(r => selectProjectSuccess(action)),
          catchError(err => of(loadProjectsError(err)))
        )
      )
    )
  );

  navigate$ = createEffect(
    () =>
      this.actions$.pipe(
        ofType(selectProjectSuccess),
        tap(() => this.router.navigate(['project']))
      ),
    { dispatch: false }
  );

  constructor(
    private actions$: Actions,
    private service: ProjectService,
    private router: Router
  ) {}
}
