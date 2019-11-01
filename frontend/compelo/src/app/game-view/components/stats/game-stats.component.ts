import { Component, Input } from '@angular/core';
import { GameStats } from '@shared/models';

@Component({
  selector: 'app-game-stats',
  template: `
    <ng-container *ngIf="gameStats.maxScoreSum as sum">
      <div class="row align-items-center mb-3">
        <div class="col">
          <h6>Most Goals</h6>
          <span>{{ sum.date | date }}: {{ sum.teams | teams }}</span>
        </div>
        <div class="col">
          <h5 class="font-weight-ligh text-center align-middle">
            {{ sum.teams | scores }}
          </h5>
        </div>
      </div>
    </ng-container>
    <ng-container *ngIf="gameStats.maxScoreDiff as lead">
      <div class="row align-items-center">
        <div class="col">
          <h6>Biggest Lead</h6>
          <span>{{ lead.date | date }}: {{ lead.teams | teams }}</span>
        </div>
        <div class="col">
          <h5 class="font-weight-ligh text-center align-middle">
            {{ lead.teams | scores }}
          </h5>
        </div>
      </div>
    </ng-container>
  `,
  styles: [
    `
      :host {
        padding: 0.75rem;
        display: block;
        background-color: white;
        border: 1px solid lightgray;
      }
    `
  ]
})
export class GameStatsComponent {
  @Input()
  gameStats: GameStats;
}
