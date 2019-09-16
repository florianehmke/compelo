import { Component, Input } from '@angular/core';
import { Player } from '@shared/models';

@Component({
  selector: 'app-player-list',
  template: `
    <app-list-group>
      <app-list-group-item *ngFor="let player of players">
        {{ player?.name }}
      </app-list-group-item>
    </app-list-group>
  `
})
export class PlayerListComponent {
  @Input()
  players: Player[];
}
