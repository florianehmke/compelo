import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { ProjectListViewComponent } from './project-list-view.component';

const routes: Routes = [
  {
    path: '',
    component: ProjectListViewComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ProjectListViewRoutingModule {}
