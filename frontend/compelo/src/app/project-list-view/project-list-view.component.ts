import { Component } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Store } from '@ngrx/store';

import { getLoadedByActionTypeOf, State as AppState } from '@core/app';
import {
  createProject,
  getProjects,
  loadProjects,
  selectProject,
  SelectProjectPayload,
  selectProjectSuccess,
  State,
} from '@core/project-list';
import { CreateProjectRequest, Project } from '@generated/api';
import { tokenForProjectIdExists } from '@shared/jwt';
import { Payload } from '@shared/models';
import { noop } from '@shared/util';

import { ProjectSelectModalComponent } from './components';

@Component({
  template: `
    <app-project-create
      (projectCreated)="onProjectCreate($event)"
    ></app-project-create>
    <hr />
    <app-project-list
      [projects]="projects$ | async"
      [isLoaded]="loaded$ | async"
      (projectSelected)="onSelect($event)"
    ></app-project-list>
  `,
})
export class ProjectListViewComponent {
  projects$ = this.store.select(getProjects);

  loaded$ = this.appStore.select(getLoadedByActionTypeOf(loadProjects));

  constructor(
    private store: Store<State>,
    private appStore: Store<AppState>,
    private modalService: NgbModal
  ) {}

  onSelect(project: Project) {
    if (tokenForProjectIdExists(project.guid)) {
      this.store.dispatch(selectProjectSuccess({ payload: project }));
    } else {
      this.modalService
        .open(ProjectSelectModalComponent)
        .result.then((pw: string) => {
          const payload: Payload<SelectProjectPayload> = {
            payload: {
              request: {
                projectGuid: project.guid,
                password: pw,
              },
              project,
            },
          };
          this.store.dispatch(selectProject(payload));
        }, noop);
    }
  }

  onProjectCreate(payload: CreateProjectRequest) {
    this.store.dispatch(createProject({ payload }));
  }
}
