import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ActionResolver } from '@core/router';
import {
  triggerLoadGames,
  triggerLoadMatches,
  triggerLoadPlayers,
  triggerLoadPlayersWithStats
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
        triggerLoadPlayers,
        triggerLoadMatches,
        triggerLoadPlayersWithStats
      ]
    }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class GameViewRoutingModule {}
