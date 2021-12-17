import { ChangeDetectionStrategy, Component } from '@angular/core';
import { getLoading, State as AppState } from '@core/app';
import { getSelectedGame, State as ProjectState } from '@core/project';
import { getSelectedProject } from '@core/project-list';
import { Store } from '@ngrx/store';

@Component({
  selector: 'app-root',
  template: `
    <app-toast></app-toast>
    {{ loading$ | async | json }}
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
  loading$ = this.appStore.select(getLoading);

  constructor(
    private projectStore: Store<ProjectState>,
    private appStore: Store<AppState>
  ) {}
}
