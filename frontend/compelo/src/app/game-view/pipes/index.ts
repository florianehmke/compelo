import { ScoresPipe } from './scores.pipe';
import { TeamPipe } from './team.pipe';
import { TeamsPipe } from './teams.pipe';

export * from './team.pipe';
export * from './teams.pipe';
export * from './scores.pipe';

export const pipes: any[] = [TeamPipe, TeamsPipe, ScoresPipe];
