import { NgFor } from '@angular/common';
import { Component, Input } from '@angular/core';

import { Player } from '@generated/api';

import { ListGroupItemComponent } from '../../../shared/list-group/list-group-item.component';
import { ListGroupComponent } from '../../../shared/list-group/list-group.component';

@Component({
  selector: 'app-player-list',
  template: `
    <app-list-group>
      <app-list-group-item *ngFor="let player of players">
        {{ player?.name }}
      </app-list-group-item>
    </app-list-group>
  `,
  standalone: true,
  imports: [ListGroupComponent, NgFor, ListGroupItemComponent],
})
export class PlayerListComponent {
  @Input()
  players: Player[];
}
