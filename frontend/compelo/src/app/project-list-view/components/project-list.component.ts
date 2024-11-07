import { NgFor } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';

import { Project } from '@generated/api';

import { ListGroupButtonComponent } from '../../../shared/list-group/list-group-button.component';
import { ListGroupComponent } from '../../../shared/list-group/list-group.component';

@Component({
  selector: 'app-project-list',
  template: `
    <app-list-group>
      <app-list-group-button
        *ngFor="let project of projects"
        (click)="projectSelected.emit(project)"
      >
        {{ project?.name }}
      </app-list-group-button>
    </app-list-group>
  `,
  standalone: true,
  imports: [ListGroupComponent, NgFor, ListGroupButtonComponent],
})
export class ProjectListComponent {
  @Input()
  projects: Project[];

  @Output()
  projectSelected = new EventEmitter<Project>();
}
