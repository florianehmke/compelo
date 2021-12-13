import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { triggerLoadProjects } from '@core/project-list';
import { ActionResolver } from '@core/router';

import { ProjectListViewComponent } from './project-list-view.component';

const routes: Routes = [
  {
    path: '',
    component: ProjectListViewComponent,
    resolve: [ActionResolver],
    data: {
      actionFactory: [triggerLoadProjects],
    },
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class ProjectListViewRoutingModule {}
