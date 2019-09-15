import { Component, EventEmitter, Input, Output } from '@angular/core';
import { Project } from '@shared/models';

@Component({
  selector: 'app-project-grid',
  template: `
    <div class="card-deck">
      <app-project-card
        *ngFor="let project of projects"
        (click)="projectSelected.emit(project)"
        [project]="project"
      ></app-project-card>
    </div>
  `
})
export class ProjectGridComponent {
  @Input()
  projects: Project[];

  @Output()
  projectSelected = new EventEmitter<Project>();
}
