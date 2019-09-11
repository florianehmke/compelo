import { NgModule } from '@angular/core';
import { SharedModuleModule } from '../../shared/shared.module';
import { ProjectViewComponent } from './project-view.component';
import { ProjectListRoutingModule } from '../project-list/project-list-routing.module';
import { ProjectViewRoutingModule } from './project-view-routing.module';

@NgModule({
  declarations: [ProjectViewComponent],
  exports: [ProjectViewComponent],
  imports: [SharedModuleModule, ProjectViewRoutingModule]
})
export class ProjectViewModule {}
