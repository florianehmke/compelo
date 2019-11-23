import { Component, Input } from '@angular/core';
import { GameStats } from '@api';

@Component({
  selector: 'app-game-stats',
  template: `
    <div class="compelo-container">
      <h6>Many Goals</h6>
      <ng-container *ngFor="let sum of gameStats.maxScoreSum">
        <div class="row align-items-center">
          <div class="col">
            <span>{{ sum.date | date }}</span>
          </div>
          <div class="col">
            <span>{{ sum.teams | teams }}</span>
          </div>
          <div class="col">
            <h5 class="font-weight-light text-center align-middle">
              {{ sum.teams | scores }}
            </h5>
          </div>
        </div>
      </ng-container>
      <h6 class="mt-3">Large Lead</h6>
      <ng-container *ngFor="let lead of gameStats.maxScoreDiff">
        <div class="row align-items-center">
          <div class="col">
            <span>{{ lead.date | date }}</span>
          </div>
          <div class="col">
            <span>{{ lead.teams | teams }}</span>
          </div>
          <div class="col">
            <h5 class="font-weight-light text-center align-middle">
              {{ lead.teams | scores }}
            </h5>
          </div>
        </div>
      </ng-container>
    </div>
  `
})
export class GameStatsComponent {
  @Input()
  gameStats: GameStats;
}
