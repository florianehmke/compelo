import { Component, Input } from '@angular/core';
import { PlayerStats, Stats } from '@shared/models';
import { StatsBarData } from '@shared/stats-bar/stats-bar.models';

@Component({
  selector: 'app-player-stats-table',
  template: `
    <div class="table-container">
      <ng-container *ngFor="let player of players">
        <div class="row mt-2" *ngIf="player.current as stats">
          <div class="col">
            <span>{{ player?.name }}</span>
          </div>
          <div class="col text-center">
            <small [ngbTooltip]="tooltips.lowestRating" class="text-muted">
              {{ stats.lowestRating }}
            </small>
            <span [ngbTooltip]="tooltips.rating">
              {{ stats.rating }}
            </span>
            <small [ngbTooltip]="tooltips.peakRating" class="text-muted">
              {{ stats.peakRating }}
            </small>
          </div>
          <div class="col d-flex justify-content-between">
            <small [ngbTooltip]="tooltips.totalGames">
              {{ stats.gameCount }}
            </small>
            <small [ngbTooltip]="tooltips.winPercentage">
              {{ winPercentage(stats) }}%
            </small>
          </div>
        </div>
        <app-stats-bar [data]="statsBarData(player?.current)"></app-stats-bar>
      </ng-container>
    </div>
  `,
  styles: [
    `
      .table-container {
        padding: 0 12px 12px 12px;
        border: 1px solid lightgray;
        background-color: white;
      }
    `
  ]
})
export class PlayerStatsTableComponent {
  @Input()
  players: PlayerStats[];

  tooltips = {
    winPercentage: 'Win Percentage',
    totalGames: 'Total Games',
    peakRating: 'Highest Elo',
    rating: 'Current Elo',
    lowestRating: 'Lowest Elo'
  };

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
