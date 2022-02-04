import { Pipe, PipeTransform } from '@angular/core';

import { MatchTeam } from '@generated/api';

@Pipe({
  name: 'teams',
})
export class TeamsPipe implements PipeTransform {
  transform(teams: MatchTeam[], ...args: any[]): string {
    if (teams) {
      return teams
        .map((team) => team.players.map((p) => p.name).join(', '))
        .join(' vs. ');
    }
    return '';
  }
}
