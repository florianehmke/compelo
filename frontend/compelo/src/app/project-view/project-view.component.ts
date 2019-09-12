import { Component } from '@angular/core';
import { Store } from '@ngrx/store';

import { State } from '../../core/project-view/project-view.reducer';
import {
  getGames,
  getSelectedGame
} from '../../core/project-view/project-view.selectors';
import {
  loadGames,
  loadPlayers
} from '../../core/project-view/project-view.actions';

@Component({
  selector: 'app-project-view',
  template: `
    <p>Project!</p>
    {{ selectedGame$ | async | json }}
    {{ games$ | async | json }}
  `
})
export class ProjectViewComponent {
  games$ = this.store.select(getGames);
  selectedGame$ = this.store.select(getSelectedGame);

  constructor(private store: Store<State>) {
    this.store.dispatch(loadGames());
    this.store.dispatch(loadPlayers());
  }
}
