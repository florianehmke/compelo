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
  loadProjects,
  selectProject,
  selectProjectSuccess,
  State
} from '@core/project-list';

@Component({
  selector: 'app-project-list',
  template: `
    <p class="lead">
      Create / Select Project
    </p>
    <app-project-create
      (projectCreated)="onProjectCreate($event)"
    ></app-project-create>
    <hr />
    <app-project-grid
      [projects]="projects$ | async"
      (projectSelected)="onSelect($event)"
    ></app-project-grid>
  `
})
export class ProjectListViewComponent {
  projects$ = this.store.select(getProjects);

  constructor(private store: Store<State>, private modalService: NgbModal) {
    this.store.dispatch(loadProjects());
  }

  onSelect(project: Project) {
    if (tokenForProjectIdExists(project.id)) {
      this.store.dispatch(selectProjectSuccess({ payload: project }));
    } else {
      this.modalService
        .open(ProjectSelectModalComponent)
        .result.then((password: string) => {
          const payload = { payload: { ...project, password: password } };
          this.store.dispatch(selectProject(payload));
        }, noop);
    }
  }

  onProjectCreate(payload: CreateProjectPayload) {
    this.store.dispatch(createProject({ payload }));
  }
}
