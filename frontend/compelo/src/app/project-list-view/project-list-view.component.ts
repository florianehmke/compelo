import { Component } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Store } from '@ngrx/store';

import {
  createProject,
  getProjects,
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
      (projectSelected)="onSelect($event)"
    ></app-project-list>
  `,
})
export class ProjectListViewComponent {
  projects$ = this.store.select(getProjects);

  constructor(private store: Store<State>, private modalService: NgbModal) {}

  onSelect(project: Project) {
    if (tokenForProjectIdExists(project.id)) {
      this.store.dispatch(selectProjectSuccess({ payload: project }));
    } else {
      this.modalService
        .open(ProjectSelectModalComponent)
        .result.then((pw: string) => {
          const payload: Payload<SelectProjectPayload> = {
            payload: {
              request: {
                projectId: project.id,
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
