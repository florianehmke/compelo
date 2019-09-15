import { Pipe, PipeTransform } from '@angular/core';
import { Team } from '@shared/models';

@Pipe({
  name: 'team'
})
export class TeamPipe implements PipeTransform {
  transform(team: Team, ...args: any[]): string {
    if (team) {
      const players = team.players.map(p => p.name).join(',');
      return `${players} (${team.score})`;
    }
    return '';
  }
}
