import { Component, EventEmitter, Input, Output } from '@angular/core';

import { MatchData } from '@generated/api';

@Component({
  selector: 'app-match-list',
  template: `
    <div class="d-flex justify-content-between">
      <p class="lead">Recent Matches</p>
      <input
        type="text"
        class="form-control-sm"
        placeholder="Filter by Player"
        (input)="filterChange.emit($event.target.value)"
      />
    </div>
    <ng-container *ngIf="isLoaded; else showLoading">
      <table class="table table-bordered bg-white">
        <tbody>
          <tr *ngFor="let match of matches | slice: from():to()">
            <td>
              <div>{{ match.date | date }}</div>
              <div>
                <small class="text-muted">{{
                  match.date | date: 'shortTime'
                }}</small>
              </div>
            </td>
            <td>
              <div class="d-flex justify-content-between">
                <div
                  *ngFor="
                    let team of match.teams;
                    let first = first;
                    let last = last
                  "
                >
                  <div [ngClass]="matchClassIf(first, last)">
                    <span>{{ team | team }}</span>
                  </div>
                  <div>
                    <small class="text-muted">Score: </small>
                    <small>{{ team?.score }}</small>
                    <small class="text-muted">, Rating: </small>
                    <small [ngClass]="ratingClassFor(team?.ratingDelta)">{{
                      team?.ratingDelta
                    }}</small>
                  </div>
                </div>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <ngb-pagination
        [(page)]="page"
        [pageSize]="pageSize"
        [maxSize]="5"
        [boundaryLinks]="false"
        [collectionSize]="matches.length"
      ></ngb-pagination>
    </ng-container>
    <ng-template #showLoading>
      <app-loading-spinner></app-loading-spinner>
    </ng-template>
  `,
  styles: [
    `
      td {
        padding: 0.5rem;
      }
    `,
  ],
})
export class MatchListComponent {
  @Input()
  matches: MatchData[];

  @Input()
  isLoaded: boolean;

  @Output()
  filterChange = new EventEmitter<string>();

  page = 1;
  pageSize = 7;

  from(): number {
    return (this.page - 1) * this.pageSize;
  }

  to(): number {
    return (this.page - 1) * this.pageSize + this.pageSize;
  }

  matchClassIf(first: boolean, last: boolean): string {
    if (first) {
      return 'text-left';
    }
    if (last) {
      return 'text-right';
    }
    return 'text-center';
  }

  ratingClassFor(delta: number): string {
    return 0 < delta ? 'text-success' : 'text-danger';
  }
}
