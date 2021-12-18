import { Component } from '@angular/core';
import { getLoadedByActionTypeOf, State as AppState } from '@core/app';
import {
  createGame,
  createPlayer,
  getGames,
  getPlayers,
  loadGames,
  loadPlayers,
  State,
} from '@core/project';
import { Game, Player } from '@generated/api';
import { Store } from '@ngrx/store';

@Component({
  template: `
    <div class="row">
      <div class="col-md-6 mb-3">
        <app-game-create
          (gameCreated)="onGameCreated($event)"
        ></app-game-create>
        <hr />
        <app-game-list
          [isLoaded]="gamesLoaded$ | async"
          [games]="games$ | async"
        ></app-game-list>
      </div>
      <div class="col-md-6">
        <app-player-create
          (playerCreated)="onPlayerCreated($event)"
        ></app-player-create>
        <hr />
        <app-player-list
          [isLoaded]="playersLoaded$ | async"
          [players]="players$ | async"
        ></app-player-list>
      </div>
    </div>
  `,
})
export class ProjectViewComponent {
  games$ = this.store.select(getGames);

  gamesLoaded$ = this.appStore.select(getLoadedByActionTypeOf(loadGames));

  players$ = this.store.select(getPlayers);

  playersLoaded$ = this.appStore.select(getLoadedByActionTypeOf(loadPlayers));

  constructor(private store: Store<State>, private appStore: Store<AppState>) {}

  onGameCreated(game: Game) {
    this.store.dispatch(createGame({ payload: game }));
  }

  onPlayerCreated(player: Player) {
    this.store.dispatch(createPlayer({ payload: player }));
  }
}
