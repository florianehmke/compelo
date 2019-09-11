import { Component } from '@angular/core';
import { Store } from '@ngrx/store';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';

import {
  getProjects,
  getSelectedProject
} from '../../core/project/project.selectors';
import { State } from '../../core/project/project.reducer';
import {
  createProject,
  loadProjects,
  selectProject,
  selectProjectSuccess
} from '../../core/project/project.actions';
import { CreateProjectPayload, Project } from '../../shared/models';
import { ProjectSelectModalComponent } from './project-select-modal.component';

@Component({
  selector: 'app-project-list',
  template: `
    <app-project-create
      (projectCreated)="onProjectCreate($event)"
    ></app-project-create>
    <hr />
    <app-project-grid
      [projects]="projects$ | async"
      [selectedProject]="selectedProject$ | async"
      (projectSelected)="onSelect($event)"
    ></app-project-grid>
  `
})
export class ProjectListComponent {
  projects$ = this.store.select(getProjects);
  selectedProject$ = this.store.select(getSelectedProject);

  constructor(private store: Store<State>, private modalService: NgbModal) {
    this.store.dispatch(loadProjects());
  }

  onSelect(project: Project) {
    const token = localStorage.getItem('compelo-token');
    if (token) {
      this.store.dispatch(selectProjectSuccess({ payload: project }));
    } else {
      this.modalService
        .open(ProjectSelectModalComponent)
        .result.then(password => {
          const payload = { payload: { ...project, password: password } };
          this.store.dispatch(selectProject(payload));
        });
    }
  }

  onProjectCreate(payload: CreateProjectPayload) {
    this.store.dispatch(createProject({ payload }));
  }
}
