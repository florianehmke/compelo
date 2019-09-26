import { Component } from '@angular/core';
import { Store } from '@ngrx/store';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Project } from '@shared/models';
import { tokenForProjectIdExists } from '@shared/jwt';
import { noop } from '@shared/util';

import { ProjectSelectModalComponent } from './components';
import {
  createProject,
  CreateProjectPayload,
  getProjects,
  selectProject,
  selectProjectSuccess,
  State
} from '@core/project-list';

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
  `
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
          const payload = {
            payload: {
              projectId: project.id,
              projectName: project.name,
              password: pw
            }
          };
          this.store.dispatch(selectProject(payload));
        }, noop);
    }
  }

  onProjectCreate(payload: CreateProjectPayload) {
    this.store.dispatch(createProject({ payload }));
  }
}
