import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ActionResolver } from '@core/router';
import { triggerLoadGames, triggerLoadPlayers } from '@core/project';
import { triggerLoadProjects } from '@core/project-list';
import { gameIdParam, projectIdParam } from '@shared/route-params';

import { ProjectViewComponent } from './project-view.component';
import { ProjectViewGuard } from './project-view.guard';

const routes: Routes = [
  {
    path: `:${projectIdParam}`,
    component: ProjectViewComponent,
    canActivate: [ProjectViewGuard],
    resolve: [ActionResolver],
    data: {
      actionFactory: [triggerLoadProjects, triggerLoadGames, triggerLoadPlayers]
    }
  },
  {
    path: `:${projectIdParam}/game/:${gameIdParam}`,
    loadChildren: () =>
      import('../game-view/game-view.module').then(mod => mod.GameViewModule)
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  providers: [ActionResolver],
  exports: [RouterModule]
})
export class ProjectViewRoutingModule {}
