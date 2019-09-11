import { NgModule } from '@angular/core';
import { SharedModuleModule } from '../../shared/shared.module';
import { ProjectViewComponent } from './project-view.component';
import { ProjectViewRoutingModule } from './project-view-routing.module';
import { ProjectViewGuard } from './project-view.guard';

@NgModule({
  declarations: [ProjectViewComponent],
  exports: [ProjectViewComponent],
  providers: [ProjectViewGuard],
  imports: [SharedModuleModule, ProjectViewRoutingModule]
})
export class ProjectViewModule {}
