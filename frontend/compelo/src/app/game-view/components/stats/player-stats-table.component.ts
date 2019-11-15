import { Component, Input } from '@angular/core';
import { PlayerStats, Stats } from '@shared/models';
import { StatsBarData } from '@shared/stats-bar/stats-bar.models';

@Component({
  selector: 'app-player-stats-table',
  template: `
    <div class="table-container">
      <ng-container *ngFor="let player of players">
        <div class="row mt-1" *ngIf="player.current as stats">
          <div class="col-4">
            {{ player?.name }}
          </div>
          <div class="col-6">
            <small class="text-muted">{{ stats.lowestRating }}</small>
            <span> {{ stats.rating }} </span>
            <small class="text-muted">{{ stats.peakRating }}</small>
          </div>
          <div class="col-2">
            <div class="d-flex justify-content-between">
              <div>
                {{ stats.gameCount }}
              </div>
              <div class="text-right">{{ winPercentage(stats) }}%</div>
            </div>
          </div>
        </div>
        <app-stats-bar [data]="statsBarData(player?.current)"></app-stats-bar>
      </ng-container>
    </div>
  `,
  styles: [
    `
      .table-container {
        padding: 0 8px 8px 8px;
        border: 1px solid lightgray;
        background-color: white;
      }
    `
  ]
})
export class PlayerStatsTableComponent {
  @Input()
  players: PlayerStats[];

  statsBarData(stats: Stats): StatsBarData {
    return {
      wins: stats.winCount,
      draws: stats.drawCount,
      lost: stats.lossCount
    };
  }

  winPercentage(stats: Stats): number {
    return Math.round((stats.winCount / stats.gameCount) * 100);
  }
}
