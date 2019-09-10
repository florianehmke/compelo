import { Injectable } from '@angular/core';
import { catchError, map, switchMap } from 'rxjs/operators';
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
          map(response => {
            localStorage.setItem('compelo-token', response.token);
            this.router.navigate(['/project']);
            return selectProjectSuccess(action);
          }),
          catchError(err => of(loadProjectsError(err)))
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
