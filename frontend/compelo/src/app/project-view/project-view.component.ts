import { Component } from '@angular/core';
import { Store } from '@ngrx/store';
import {
  createGame,
  createPlayer,
  getGames,
  getPlayers,
  getSelectedGame,
  loadGames,
  loadPlayers,
  selectGame,
  State
} from '@core/project';
import { Game, Player } from '@shared/models';

@Component({
  selector: 'app-project-view',
  template: `
    <div class="row">
      <div class="col">
        <p class="lead">
          Create / Select Game
        </p>
        <app-game-create
          (gameCreated)="onGameCreated($event)"
        ></app-game-create>
        <hr />
        <div class="list-group">
          <button
            [routerLink]="['game', game.id]"
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
        <p class="lead">
          Create Player
        </p>
        <app-player-create
          (playerCreated)="onPlayerCreated($event)"
        ></app-player-create>
        <hr />
        <ul class="list-group">
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
