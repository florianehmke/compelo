import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ActionResolver } from '@core/router';
import {
  triggerLoadGames,
  triggerLoadMatches,
  triggerLoadPlayers
} from '@core/project';

import { GameViewComponent } from './game-view.component';

const routes: Routes = [
  {
    path: '',
    component: GameViewComponent,
    resolve: [ActionResolver],
    data: {
      actionFactory: [triggerLoadGames, triggerLoadPlayers, triggerLoadMatches]
    }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class GameViewRoutingModule {}
