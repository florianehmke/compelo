import { NgFor } from '@angular/common';
import { Component, Input } from '@angular/core';
import { RouterLink } from '@angular/router';

import { Game } from '@generated/api';

import { ListGroupButtonComponent } from '../../../shared/list-group/list-group-button.component';
import { ListGroupComponent } from '../../../shared/list-group/list-group.component';


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
  standalone: true,
  imports: [ListGroupComponent, NgFor, ListGroupButtonComponent, RouterLink],
})
export class GameListComponent {
  @Input()
  games: Game[];
}
