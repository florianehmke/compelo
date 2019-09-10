import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ProjectListComponent } from './project-list.component';
import { ProjectComponent } from './project.component';

const routes: Routes = [
  {
    path: 'list',
    component: ProjectListComponent
  },
  {
    path: 'project',
    component: ProjectComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ProjectRoutingModule {}
