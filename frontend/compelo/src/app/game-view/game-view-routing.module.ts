import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import {
  triggerLoadCompetitions,
  triggerLoadGames,
  triggerLoadGameStats,
  triggerLoadMatches,
  triggerLoadPlayers,
  triggerLoadPlayerStats,
} from '@core/project';
import { triggerLoadProjects } from '@core/project-list';
import { ActionResolver } from '@core/router';

import { GameViewComponent } from './game-view.component';

const routes: Routes = [
  {
    path: '',
    component: GameViewComponent,
    resolve: [ActionResolver],
    data: {
      actionFactory: [
        triggerLoadProjects,
        triggerLoadGames,
        triggerLoadGameStats,
        triggerLoadPlayers,
        triggerLoadPlayerStats,
        triggerLoadMatches,
        triggerLoadCompetitions,
      ],
    },
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class GameViewRoutingModule {}
