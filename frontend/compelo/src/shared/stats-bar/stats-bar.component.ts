import { Component, Input } from '@angular/core';

import { StatsBarData } from '@shared/stats-bar/stats-bar.models';

@Component({
  selector: 'app-stats-bar',
  template: `
    <div class="bar-container" *ngIf="data">
      <div class="segment wins" [style.flex]="data.wins">
        <span *ngIf="data.wins > 0">{{ data.wins }}</span>
      </div>
      <div class="segment draws" [style.flex]="data.draws">
        <span *ngIf="data.draws > 0">{{ data.draws }}</span>
      </div>
      <div class="segment lost" [style.flex]="data.lost">
        <span *ngIf="data.lost > 0">{{ data.lost }}</span>
      </div>
    </div>
  `,
  styles: [
    `
      .bar-container {
        height: 1.2rem;
        font-size: x-small;
        color: black;
        width: 100%;
        display: flex;
        flex-direction: row;
      }

      .segment {
        display: flex;
        justify-content: center;
        align-items: center;
      }

      .wins {
        background-color: lightgreen;
      }

      .draws {
        background-color: yellow;
      }

      .lost {
        background-color: lightcoral;
      }
    `,
  ],
})
export class StatsBarComponent {
  @Input() data: StatsBarData;
}
