import { MatchCreateComponent } from './match-create.component';
import { MatchSettingsComponent } from './match-settings.component';
import { MatchListComponent } from './match-list.component';
import { LeaderboardComponent } from './leaderboard.component';

export * from './match-create.component';
export * from './match-settings.component';
export * from './match-list.component';
export * from './leaderboard.component';

export const components: any[] = [
  MatchSettingsComponent,
  MatchCreateComponent,
  MatchListComponent,
  LeaderboardComponent
];
