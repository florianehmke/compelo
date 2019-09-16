import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ActionResolver } from '@core/router';
import { triggerLoadGamesAndMatches, triggerLoadPlayers } from '@core/project';

import { GameViewComponent } from './game-view.component';

const routes: Routes = [
  {
    path: '',
    component: GameViewComponent,
    resolve: [ActionResolver],
    data: {
      actionFactory: [triggerLoadGamesAndMatches, triggerLoadPlayers]
    }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class GameViewRoutingModule {}
