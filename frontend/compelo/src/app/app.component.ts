import { ChangeDetectionStrategy, Component } from '@angular/core';
import { getSelectedGame, State as ProjectState } from '@core/project';

import { Store } from '@ngrx/store';
import { getSelectedProject } from '@core/project-list';

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
})
export class AppComponent {
  game$ = this.projectStore.select(getSelectedGame);
  project$ = this.projectStore.select(getSelectedProject);

  constructor(private projectStore: Store<ProjectState>) {}
}
