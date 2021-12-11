import { Component, Input } from '@angular/core';
import { Game } from '@api';

@Component({
  selector: 'app-game-list',
  template: `
    <app-list-group>
      <app-list-group-button
        *ngFor="let game of games"
        [routerLink]="['game', game.id]"
      >
        {{ game?.name }}
      </app-list-group-button>
    </app-list-group>
  `,
})
export class GameListComponent {
  @Input()
  games: Game[];
}
