import { Component, Input } from '@angular/core';
import { Game } from '@generated/api';

@Component({
  selector: 'app-game-list',
  template: `
    <app-list-group *ngIf="isLoaded; else showLoading">
      <app-list-group-button
        *ngFor="let game of games"
        [routerLink]="['game', game.id]"
      >
        {{ game?.name }}
      </app-list-group-button>
    </app-list-group>
    <ng-template #showLoading>
      <app-loading-spinner></app-loading-spinner>
    </ng-template>
  `,
})
export class GameListComponent {
  @Input()
  games: Game[];

  @Input()
  isLoaded: boolean;
}
