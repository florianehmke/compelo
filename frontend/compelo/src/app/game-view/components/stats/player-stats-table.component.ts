import { Component, Input } from '@angular/core';
import { PlayerStats } from '@shared/models';

@Component({
  selector: 'app-player-stats-table',
  template: `
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
            <small class="text-muted">{{
              player?.current?.lowestRating
            }}</small>
            <span> {{ player?.current?.rating }} </span>
            <small class="text-muted">{{ player?.current?.peakRating }}</small>
          </td>
          <td class="text-center">{{ player?.current?.winCount }}</td>
          <td class="text-center">{{ player?.current?.drawCount }}</td>
          <td class="text-center">{{ player?.current?.lossCount }}</td>
        </tr>
      </tbody>
    </table>
  `
})
export class PlayerStatsTableComponent {
  @Input()
  players: PlayerStats[];
}
