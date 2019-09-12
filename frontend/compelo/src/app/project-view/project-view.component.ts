import { Component } from '@angular/core';
import { Store } from '@ngrx/store';

import { State } from '../../core/project-view/project-view.reducer';
import {
  getGames,
  getPlayers,
  getSelectedGame
} from '../../core/project-view/project-view.selectors';
import {
  createGame,
  createPlayer,
  loadGames,
  loadPlayers,
  selectGame
} from '../../core/project-view/project-view.actions';
import { Game, Player } from '../../shared/models';

@Component({
  selector: 'app-project-view',
  template: `
    <div class="row">
      <div class="col">
        <app-game-create
          (gameCreated)="onGameCreated($event)"
        ></app-game-create>

        <div class="list-group mt-4">
          <button
            *ngFor="let game of games$ | async"
            type="button"
            class="list-group-item list-group-item-action"
            (click)="selectGame(game)"
          >
            {{ game?.name }}
          </button>
        </div>
      </div>
      <div class="col">
        <app-player-create
          (playerCreated)="onPlayerCreated($event)"
        ></app-player-create>

        <ul class="list-group mt-4">
          <li
            *ngFor="let player of players$ | async"
            class="list-group-item list-group-item-light"
          >
            {{ player?.name }}
          </li>
        </ul>
      </div>
    </div>
  `
})
export class ProjectViewComponent {
  games$ = this.store.select(getGames);
  selectedGame$ = this.store.select(getSelectedGame);
  players$ = this.store.select(getPlayers);

  constructor(private store: Store<State>) {
    this.store.dispatch(loadGames());
    this.store.dispatch(loadPlayers());
  }

  onGameCreated(game: Game) {
    this.store.dispatch(createGame({ payload: game }));
  }

  onPlayerCreated(player: Player) {
    this.store.dispatch(createPlayer({ payload: player }));
  }

  selectGame(game: Game) {
    this.store.dispatch(selectGame({ payload: game }));
  }
}
