import { Pipe, PipeTransform } from '@angular/core';

import { MatchTeam } from '@generated/api';

@Pipe({
  name: 'scores',
})
export class ScoresPipe implements PipeTransform {
  transform(teams: MatchTeam[], ...args: any[]): string {
    if (teams) {
      return teams.map((team) => team.score).join(':');
    }
    return '';
  }
}
