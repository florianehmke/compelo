import { Pipe, PipeTransform } from '@angular/core';
import { Team } from '@shared/models';

@Pipe({
  name: 'teams'
})
export class TeamsPipe implements PipeTransform {
  transform(teams: Team[], ...args: any[]): string {
    if (teams) {
      return teams
        .map(team => team.players.map(p => p.name).join(', '))
        .join(' vs. ');
    }
    return '';
  }
}
