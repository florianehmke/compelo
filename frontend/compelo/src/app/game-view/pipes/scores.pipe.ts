import { Pipe, PipeTransform } from '@angular/core';
import { Team } from '@shared/models';

@Pipe({
  name: 'scores'
})
export class ScoresPipe implements PipeTransform {
  transform(teams: Team[], ...args: any[]): string {
    if (teams) {
      return teams.map(team => team.score).join(':');
    }
    return '';
  }
}
