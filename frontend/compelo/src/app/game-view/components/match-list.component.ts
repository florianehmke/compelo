import { Component, Input } from '@angular/core';
import { Match } from '@shared/models';

@Component({
  selector: 'app-match-list',
  template: `
    <p class="lead">Recent Matches</p>
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
      [collectionSize]="matches.length"
    ></ngb-pagination>
  `,
  styles: [
    `
      td {
        padding: 0.5rem;
      }
    `
  ]
})
export class MatchListComponent {
  @Input()
  matches: Match[];

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