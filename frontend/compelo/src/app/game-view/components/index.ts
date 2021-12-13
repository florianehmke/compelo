import { MatchCreateComponent } from './match-create.component';
import { MatchListComponent } from './match-list.component';
import { MatchSettingsComponent } from './match-settings.component';
import { GameStatsComponent } from './stats/game-stats.component';
import { PlayerStatsChartComponent } from './stats/player-stats-chart.component';
import { PlayerStatsTableComponent } from './stats/player-stats-table.component';
import { StatsComponent } from './stats/stats.component';

export * from './match-create.component';
export * from './match-settings.component';
export * from './match-list.component';

export * from './stats/stats.component';
export * from './stats/game-stats.component';
export * from './stats/player-stats-table.component';
export * from './stats/player-stats-table.component';

export const components: any[] = [
  MatchSettingsComponent,
  MatchCreateComponent,
  MatchListComponent,
  StatsComponent,
  GameStatsComponent,
  PlayerStatsChartComponent,
  PlayerStatsTableComponent,
];
