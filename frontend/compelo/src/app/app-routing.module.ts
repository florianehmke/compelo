import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: '',
    redirectTo: 'project-list',
    pathMatch: 'full'
  },
  {
    path: 'project-list',
    loadChildren: () =>
      import('./project-list-view/project-list-view.module').then(
        mod => mod.ProjectListViewModule
      )
  },
  {
    path: 'project-view',
    loadChildren: () =>
      import('./project-view/project-view.module').then(
        mod => mod.ProjectViewModule
      )
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
