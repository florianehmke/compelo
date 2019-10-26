import { Component, Input } from '@angular/core';
import { PlayerStats } from '@shared/models';
import { IconName } from '@fortawesome/fontawesome-svg-core';

@Component({
  selector: 'app-player-stats',
  template: `
    <p class="lead">
      Player Stats
      <app-icon
        class="float-right"
        [icon]="icon"
        [button]="true"
        (click)="toggle()"
      ></app-icon>
    </p>
    <ng-container *ngIf="showTable">
      <app-player-stats-table [players]="players"></app-player-stats-table>
    </ng-container>
    <ng-container *ngIf="showChart">
      <app-player-stats-chart [players]="players"></app-player-stats-chart>
    </ng-container>
  `
})
export class PlayerStatsComponent {
  @Input()
  players: PlayerStats[];

  show: 'table' | 'chart' = 'table';

  toggle() {
    this.show = this.show == 'table' ? 'chart' : 'table';
  }

  get icon(): IconName {
    return this.show == 'table' ? 'chart-line' : 'table';
  }

  get showChart(): boolean {
    return this.show == 'chart';
  }

  get showTable(): boolean {
    return this.show == 'table';
  }
}
