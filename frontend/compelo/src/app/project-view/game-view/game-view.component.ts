import { Component } from '@angular/core';
import { getPlayers, getSelectedGame, State } from '@core/project';
import { Store } from '@ngrx/store';

@Component({
  selector: 'app-game-view',
  template: `
    Game
  `
})
export class GameViewComponent {
  selectedGame$ = this.store.select(getSelectedGame);
  players$ = this.store.select(getPlayers);

  constructor(private store: Store<State>) {}
}
