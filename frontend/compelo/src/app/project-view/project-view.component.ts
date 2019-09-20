import { Component } from '@angular/core';
import { Store } from '@ngrx/store';
import {
  createGame,
  createPlayer,
  getGames,
  getPlayers,
  State
} from '@core/project';
import { Game, Player } from '@shared/models';

@Component({
  template: `
    <div class="row">
      <div class="col">
        <app-game-create
          (gameCreated)="onGameCreated($event)"
        ></app-game-create>
        <hr />
        <app-game-list [games]="games$ | async"></app-game-list>
      </div>
      <div class="col">
        <app-player-create
          (playerCreated)="onPlayerCreated($event)"
        ></app-player-create>
        <hr />
        <app-player-list [players]="players$ | async"></app-player-list>
      </div>
    </div>
  `
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
