import { Component, Input } from '@angular/core';
import { GameStats, PlayerStats } from '@shared/models';

@Component({
  selector: 'app-stats',
  template: `
    <p class="lead">
      Statistics
      <app-icon
        class="float-right"
        [class.text-muted]="show === 'table'"
        icon="table"
        [button]="true"
        (click)="show = 'table'"
      ></app-icon>
      <app-icon
        class="float-right"
        [class.text-muted]="show === 'chart'"
        icon="chart-line"
        [button]="true"
        (click)="show = 'chart'"
      ></app-icon>
      <app-icon
        class="float-right"
        [class.text-muted]="show === 'game'"
        icon="trophy"
        [button]="true"
        (click)="show = 'game'"
      ></app-icon>
    </p>
    <ng-container *ngIf="show === 'table'">
      <app-player-stats-table [players]="players"></app-player-stats-table>
    </ng-container>
    <ng-container *ngIf="show === 'chart'">
      <app-player-stats-chart [players]="players"></app-player-stats-chart>
    </ng-container>
    <ng-container *ngIf="show === 'game'">
      <app-game-stats [gameStats]="gameStats"></app-game-stats>
    </ng-container>
  `,
  styles: [
    `
      app-icon {
        margin-left: 0.5rem;
      }
    `
  ]
})
export class StatsComponent {
  @Input()
  players: PlayerStats[];

  @Input()
  gameStats: GameStats;

  show: 'table' | 'chart' | 'game' = 'game';
}
