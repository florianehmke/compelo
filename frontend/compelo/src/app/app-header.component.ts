import { Component, Input, OnInit } from '@angular/core';
import { Game, Project } from '@shared/models';

@Component({
  selector: 'app-header',
  template: `
    <div class="header lead">
      <a [routerLink]="['/']">
        > compelo
      </a>
      <ng-container *ngIf="project">
        <span>> </span>
        <a [routerLink]="['project-view', project?.id]">
          {{ project?.name }}
        </a>
      </ng-container>
      <ng-container *ngIf="game && project">
        <span>> </span>
        <a [routerLink]="['project-view', project?.id, 'game', game?.id]">
          {{ game?.name }}
        </a>
      </ng-container>
    </div>
  `,
  styles: [
    `
      .header {
        height: 48px;
        padding: 8px;
        margin-bottom: 32px;
      }
      a, span {
        color: white;
      }
    `
  ]
})
export class AppHeaderComponent {
  @Input()
  game: Game;

  @Input()
  project: Project;
}
