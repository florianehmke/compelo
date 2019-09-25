import { Component, Input } from '@angular/core';
import { Player } from '@shared/models';

@Component({
  selector: 'app-leaderboard',
  template: `
    <p class="lead">Leaderboard</p>
    <table class="table table-bordered bg-white">
      <thead>
        <tr>
          <th scope="col">Name</th>
          <th scope="col" class="text-center">Rating</th>
          <th scope="col" class="text-center">W</th>
          <th scope="col" class="text-center">D</th>
          <th scope="col" class="text-center">L</th>
        </tr>
      </thead>
      <tbody>
        <tr *ngFor="let player of players">
          <td>{{ player?.name }}</td>
          <td class="text-center">
            <small class="text-muted">{{ player?.lowestRating }}</small>
            <span> {{ player?.rating }} </span>
            <small class="text-muted">{{ player?.peakRating }}</small>
          </td>
          <td class="text-center">{{ player?.winCount }}</td>
          <td class="text-center">{{ player?.drawCount }}</td>
          <td class="text-center">{{ player?.lossCount }}</td>
        </tr>
      </tbody>
    </table>
  `
})
export class LeaderboardComponent {
  @Input()
  players: Player[];
}
