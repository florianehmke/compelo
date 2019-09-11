import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { ProjectViewComponent } from './project-view.component';
import { ProjectViewGuard } from './project-view.guard';

const routes: Routes = [
  {
    path: ':id',
    component: ProjectViewComponent,
    canActivate: [ProjectViewGuard]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ProjectViewRoutingModule {}
