import { Component, EventEmitter, Input, Output } from '@angular/core';
import { Project } from '@shared/models';

@Component({
  selector: 'app-project-list',
  template: `
    <div class="list-group">
      <button
        type="button"
        class="list-group-item list-group-item-action"
        *ngFor="let project of projects"
        (click)="projectSelected.emit(project)"
      >
        {{ project?.name }}
      </button>
    </div>
  `
})
export class ProjectListComponent {
  @Input()
  projects: Project[];

  @Output()
  projectSelected = new EventEmitter<Project>();
}
