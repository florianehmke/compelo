import { NgModule } from '@angular/core';
import { SharedModule } from '@shared/shared.module';

import { ProjectViewComponent } from './project-view.component';
import { ProjectViewRoutingModule } from './project-view-routing.module';
import { ProjectViewGuard } from './project-view.guard';
import { components } from './components';

@NgModule({
  declarations: [ProjectViewComponent, ...components],
  exports: [],
  providers: [ProjectViewGuard],
  imports: [SharedModule, ProjectViewRoutingModule],
})
export class ProjectViewModule {}
