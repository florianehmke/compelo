import { Component } from '@angular/core';
import { Game, Player } from '@api';
import { Store } from '@ngrx/store';

import {
  createGame,
  createPlayer,
  getGames,
  getPlayers,
  State,
} from '@core/project';

@Component({
  template: `
    <div class="row">
      <div class="col-md-6 mb-3">
        <app-game-create
          (gameCreated)="onGameCreated($event)"
        ></app-game-create>
        <hr />
        <app-game-list [games]="games$ | async"></app-game-list>
      </div>
      <div class="col-md-6">
        <app-player-create
          (playerCreated)="onPlayerCreated($event)"
        ></app-player-create>
        <hr />
        <app-player-list [players]="players$ | async"></app-player-list>
      </div>
    </div>
  `,
})
export class ProjectViewComponent {
  games$ = this.store.select(getGames);
  players$ = this.store.select(getPlayers);

  constructor(private store: Store<State>) {}

  onGameCreated(game: Game) {
    this.store.dispatch(createGame({ payload: game }));
  }

  onPlayerCreated(player: Player) {
    this.store.dispatch(createPlayer({ payload: player }));
  }
}
