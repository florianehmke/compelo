import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ActionResolver } from '@core/router';
import {
  triggerLoadGames,
  triggerLoadGameStats,
  triggerLoadMatches,
  triggerLoadPlayers,
  triggerLoadPlayerStats
} from '@core/project';
import { triggerLoadProjects } from '@core/project-list';

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
        triggerLoadMatches
      ]
    }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class GameViewRoutingModule {}
