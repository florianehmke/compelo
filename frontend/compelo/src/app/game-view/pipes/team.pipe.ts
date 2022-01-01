import { Pipe, PipeTransform } from '@angular/core';

import { Team } from '@generated/api';

@Pipe({
  name: 'team',
})
export class TeamPipe implements PipeTransform {
  transform(team: Team, ...args: any[]): string {
    if (team) {
      return team.players.map((p) => p.name).join(', ');
    }
    return '';
  }
}
