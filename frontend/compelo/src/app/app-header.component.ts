import { NgIf } from '@angular/common';
import { Component, Input } from '@angular/core';
import { RouterLink } from '@angular/router';

import { Game, Project } from '@generated/api';

import { IconComponent } from '../shared/icon/icon.component';

@Component({
  selector: 'app-header',
  template: `
    <div class="header lead d-flex justify-content-between">
      <div>
        <a [routerLink]="['/']"> > compelo </a>
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
      <div class="mr-3">
        <a
          target="_blank"
          rel="noopener noreferrer"
          href="https://github.com/florianehmke/compelo"
        >
          <app-icon button="true" icon="github" prefix="fab"></app-icon>
        </a>
      </div>
    </div>
  `,
  styles: [
    `
      .header {
        height: 48px;
        padding: 8px;
        margin-bottom: 32px;
      }
      a,
      span {
        color: white;
      }
    `,
  ],
  standalone: true,
  imports: [RouterLink, NgIf, IconComponent],
})
export class AppHeaderComponent {
  @Input()
  game: Game;

  @Input()
  project: Project;
}
