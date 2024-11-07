import { AsyncPipe } from '@angular/common';
import { ChangeDetectionStrategy, Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { Store } from '@ngrx/store';

import { getSelectedGame, State as ProjectState } from '@core/project';
import { getSelectedProject } from '@core/project-list';

import { AppFooterComponent } from './app-footer.component';
import { AppHeaderComponent } from './app-header.component';

@Component({
  selector: 'app-root',
  template: `
    <app-toast></app-toast>
    <app-header
      [game]="game$ | async"
      [project]="project$ | async"
    ></app-header>
    <div class="container flex-grow-1">
      <router-outlet></router-outlet>
    </div>
    <app-footer></app-footer>
  `,
  changeDetection: ChangeDetectionStrategy.OnPush,
  styles: [
    `
      :host {
        display: flex;
        flex-direction: column;
        min-height: 100vh;
      }
    `,
  ],
  standalone: true,
  imports: [AppHeaderComponent, RouterOutlet, AppFooterComponent, AsyncPipe],
})
export class AppComponent {
  game$ = this.projectStore.select(getSelectedGame);
  project$ = this.projectStore.select(getSelectedProject);

  constructor(private projectStore: Store<ProjectState>) {}
}
