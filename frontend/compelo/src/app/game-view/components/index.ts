import { MatchCreateComponent } from './match-create.component';
import { MatchSettingsComponent } from './match-settings.component';
import { MatchListComponent } from './match-list.component';
import { PlayerStatsComponent } from './stats/player-stats.component';
import { PlayerStatsChartComponent } from './stats/player-stats-chart.component';
import { PlayerStatsTableComponent } from './stats/player-stats-table.component';

export * from './match-create.component';
export * from './match-settings.component';
export * from './match-list.component';

export * from './stats/player-stats.component';
export * from './stats/player-stats-table.component';
export * from './stats/player-stats-table.component';

export const components: any[] = [
  MatchSettingsComponent,
  MatchCreateComponent,
  MatchListComponent,
  PlayerStatsComponent,
  PlayerStatsChartComponent,
  PlayerStatsTableComponent
];
