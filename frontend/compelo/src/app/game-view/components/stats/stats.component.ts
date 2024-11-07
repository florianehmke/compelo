import { NgFor, NgIf } from '@angular/common';
import { Component, Input } from '@angular/core';
import { IconName } from '@fortawesome/fontawesome-svg-core';

import { GameStats, PlayerStats } from '@generated/api';

import { GameStatsComponent } from './game-stats.component';
import { PlayerStatsChartComponent } from './player-stats-chart.component';
import { PlayerStatsTableComponent } from './player-stats-table.component';
import { IconComponent } from '../../../../shared/icon/icon.component';

export interface Mode {
  title: string;
  icon: IconName;
}

@Component({
  selector: 'app-stats',
  template: `
    <p class="lead">
      {{ currentMode.title }}
      <ng-container *ngFor="let mode of [table, chart, game]">
        <app-icon
          class="float-right"
          [icon]="mode.icon"
          [button]="true"
          (click)="currentMode = mode"
        ></app-icon>
      </ng-container>
    </p>
    <ng-container *ngIf="currentMode === table">
      <app-player-stats-table [players]="players"></app-player-stats-table>
    </ng-container>
    <ng-container *ngIf="currentMode === chart">
      <app-player-stats-chart [players]="players"></app-player-stats-chart>
    </ng-container>
    <ng-container *ngIf="currentMode === game">
      <app-game-stats [gameStats]="gameStats"></app-game-stats>
    </ng-container>
  `,
  styles: [
    `
      app-icon {
        margin-left: 0.5rem;
      }
    `,
  ],
  standalone: true,
  imports: [
    NgFor,
    IconComponent,
    NgIf,
    PlayerStatsTableComponent,
    PlayerStatsChartComponent,
    GameStatsComponent,
  ],
})
export class StatsComponent {
  @Input()
  players: PlayerStats[];

  @Input()
  gameStats: GameStats;

  readonly table: Mode = {
    icon: 'table',
    title: 'Leaderboard',
  };

  readonly chart: Mode = {
    icon: 'chart-line',
    title: 'History Chart',
  };

  readonly game: Mode = {
    icon: 'trophy',
    title: 'Remarkable Games',
  };

  currentMode: Mode = this.table;
}
