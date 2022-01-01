import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { of } from 'rxjs';
import { catchError, map, switchMap, tap } from 'rxjs/operators';

import { Project } from '@generated/api';
import { storeToken } from '@shared/jwt';

import { AuthService } from '../auth.service';
import {
  createProject,
  createProjectError,
  createProjectSuccess,
  loadProjects,
  loadProjectsError,
  loadProjectsSuccess,
  selectProject,
  selectProjectSuccess,
} from './project-list.actions';
import { ProjectListService } from './project-list.service';

@Injectable()
export class ProjectListEffects {
  loadProjects$ = createEffect(() =>
    this.actions$.pipe(
      ofType(loadProjects),
      switchMap(() =>
        this.service.getProjects().pipe(
          map((projects) => loadProjectsSuccess({ payload: projects })),
          catchError((err) => of(loadProjectsError(err)))
        )
      )
    )
  );

  selectProject$ = createEffect(() =>
    this.actions$.pipe(
      ofType(selectProject),
      switchMap((action) =>
        this.authService.login(action.payload.request).pipe(
          tap((response) => storeToken(response.token)),
          map(() =>
            selectProjectSuccess({
              payload: { ...action.payload.project },
            })
          ),
          catchError((err) => of(loadProjectsError(err)))
        )
      )
    )
  );

  createProject$ = createEffect(() =>
    this.actions$.pipe(
      ofType(createProject),
      switchMap((action) =>
        this.service.createProject(action.payload).pipe(
          switchMap((response) => [
            createProjectSuccess({ payload: response }),
            selectProject({
              payload: {
                project: {
                  guid: response.guid,
                  name: action.payload.name,
                } as Project,
                request: {
                  password: action.payload.password,
                  projectGuid: response.guid,
                },
              },
            }),
          ]),
          catchError((err) => of(createProjectError(err)))
        )
      )
    )
  );

  navigate$ = createEffect(
    () =>
      this.actions$.pipe(
        ofType(selectProjectSuccess),
        tap((action) =>
          this.router.navigate(['project-view', action.payload.guid])
        )
      ),
    { dispatch: false }
  );

  constructor(
    private actions$: Actions,
    private service: ProjectListService,
    private authService: AuthService,
    private router: Router
  ) {}
}
