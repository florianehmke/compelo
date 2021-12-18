import { Component, EventEmitter, Input, Output } from '@angular/core';

import { Project } from '@generated/api';

@Component({
  selector: 'app-project-list',
  template: `
    <app-list-group *ngIf="isLoaded; else showLoading">
      <app-list-group-button
        *ngFor="let project of projects"
        (click)="projectSelected.emit(project)"
      >
        {{ project?.name }}
      </app-list-group-button>
    </app-list-group>
    <ng-template #showLoading>
      <app-loading-spinner></app-loading-spinner>
    </ng-template>
  `,
})
export class ProjectListComponent {
  @Input()
  projects: Project[];

  @Input()
  isLoaded: boolean;

  @Output()
  projectSelected = new EventEmitter<Project>();
}
