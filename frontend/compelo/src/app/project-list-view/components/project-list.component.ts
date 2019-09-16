import { Component, EventEmitter, Input, Output } from '@angular/core';
import { Project } from '@shared/models';

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
  `
})
export class ProjectListComponent {
  @Input()
  projects: Project[];

  @Output()
  projectSelected = new EventEmitter<Project>();
}
