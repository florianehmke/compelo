import { Component, Input } from '@angular/core';
import { GameStats } from '@shared/models';

@Component({
  selector: 'app-game-stats',
  template: `
    {{ gameStats | json }}
  `
})
export class GameStatsComponent {
  @Input()
  gameStats: GameStats;
}
