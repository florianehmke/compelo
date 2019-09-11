import { Component } from '@angular/core';
import { Store } from '@ngrx/store';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';

import { getProjects } from '../../core/project/project.selectors';
import { State } from '../../core/project/project.reducer';
import {
  loadProjects,
  selectProject
} from '../../core/project/project.actions';
import { Project } from '../../shared/models';
import { ProjectSelectModalComponent } from './project-select-modal.component';

@Component({
  selector: 'app-project-list',
  template: `
    <div *ngFor="let project of projects$ | async">
      <span (click)="onSelect(project)">
        {{ project.id }} {{ project.name }}
      </span>
    </div>
  `
})
export class ProjectListComponent {
  projects$ = this.store.select(getProjects);

  constructor(private store: Store<State>, private modalService: NgbModal) {
    this.store.dispatch(loadProjects());
  }

  onSelect(project: Project) {
    this.modalService
      .open(ProjectSelectModalComponent)
      .result.then(password => {
        const payload = { payload: { ...project, password: password } };
        this.store.dispatch(selectProject(payload));
      });
  }
}
