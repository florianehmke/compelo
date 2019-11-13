import { Component, Input } from '@angular/core';
import { GameStats } from '@shared/models';

@Component({
  selector: 'app-game-stats',
  template: `
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
          <h5 class="font-weight-ligh text-center align-middle">
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
