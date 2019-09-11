import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { ProjectViewComponent } from './project-view.component';

const routes: Routes = [
  {
    path: '',
    component: ProjectViewComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ProjectViewRoutingModule {}
