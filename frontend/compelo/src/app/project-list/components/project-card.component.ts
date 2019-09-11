import { Component, Input } from '@angular/core';
import { Project } from '../../../shared/models';

@Component({
  selector: 'app-project-card',
  template: `
    <div class="card" [class.border-black]="selected">
      <div class="card-body">
        {{ project.name }}
      </div>
    </div>
  `,
  styles: [
    `
      .card:hover {
        border-color: black;
        cursor: pointer;
      }
    `
  ]
})
export class ProjectCardComponent {
  @Input()
  project: Project;

  @Input()
  selected = false;
}
