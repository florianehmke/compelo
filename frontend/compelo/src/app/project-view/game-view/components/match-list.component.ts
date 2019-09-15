import { Component, Input } from '@angular/core';
import { Match } from '@shared/models';

@Component({
  selector: 'app-match-list',
  template: `
    <table class="table table-striped">
      <thead>
        <tr>
          <th scope="col">Date</th>
          <th scope="col">Match</th>
        </tr>
      </thead>
      <tbody>
        <tr *ngFor="let match of matches; index as i">
          <td>{{ match.date | date }}</td>
          <td>
            <div class="d-flex justify-content-between">
              <span *ngFor="let team of match.teams">{{ team | team }}</span>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
  `
})
export class MatchListComponent {
  @Input()
  matches: Match[];
}
