import { Component, Input } from '@angular/core';

import { Player } from '@generated/api';

@Component({
  selector: 'app-player-list',
  template: `
    <app-list-group *ngIf="isLoaded; else showLoading">
      <app-list-group-item *ngFor="let player of players">
        {{ player?.name }}
      </app-list-group-item>
    </app-list-group>
    <ng-template #showLoading>
      <app-loading-spinner></app-loading-spinner>
    </ng-template>
  `,
})
export class PlayerListComponent {
  @Input()
  players: Player[];

  @Input()
  isLoaded: boolean;
}
