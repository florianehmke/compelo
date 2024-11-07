import { NgModule } from '@angular/core';

import { SharedModule } from '@shared/shared.module';

import { components } from './components';
import { ProjectViewRoutingModule } from './project-view-routing.module';
import { ProjectViewComponent } from './project-view.component';
import { ProjectViewGuard } from './project-view.guard';

@NgModule({
  exports: [],
  providers: [ProjectViewGuard],
  imports: [
    SharedModule,
    ProjectViewRoutingModule,
    ProjectViewComponent,
    ...components,
  ],
})
export class ProjectViewModule {}
