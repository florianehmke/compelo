import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { ProjectViewComponent } from './project-view.component';
import { ProjectViewGuard } from './project-view.guard';

const routes: Routes = [
  {
    path: ':projectId',
    component: ProjectViewComponent,
    canActivate: [ProjectViewGuard]
  },
  {
    path: ':projectId/game/:gameId',
    loadChildren: () =>
      import('./game-view/game-view.module').then(mod => mod.GameViewModule)
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ProjectViewRoutingModule {}
