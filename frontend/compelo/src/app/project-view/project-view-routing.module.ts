import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ActionResolver } from '@core/router';
import { triggerLoadGames, triggerLoadPlayers } from '@core/project';

import { ProjectViewComponent } from './project-view.component';
import { ProjectViewGuard } from './project-view.guard';

const routes: Routes = [
  {
    path: ':projectId',
    component: ProjectViewComponent,
    canActivate: [ProjectViewGuard],
    resolve: [ActionResolver],
    data: {
      actionFactory: [triggerLoadGames, triggerLoadPlayers]
    }
  },
  {
    path: ':projectId/game/:gameId',
    loadChildren: () =>
      import('./game-view/game-view.module').then(mod => mod.GameViewModule)
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  providers: [ActionResolver],
  exports: [RouterModule]
})
export class ProjectViewRoutingModule {}
