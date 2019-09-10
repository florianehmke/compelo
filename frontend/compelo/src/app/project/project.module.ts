import { NgModule } from '@angular/core';

import { ProjectListComponent } from './project-list.component';
import { ProjectSelectModalComponent } from './project-select-modal.component';
import { SharedModuleModule } from '../../shared/shared.module';
import { ProjectComponent } from './project.component';
import { ProjectRoutingModule } from './project-routing.module';

@NgModule({
  declarations: [
    ProjectComponent,
    ProjectListComponent,
    ProjectSelectModalComponent
  ],
  exports: [
    ProjectComponent,
    ProjectListComponent,
    ProjectSelectModalComponent
  ],
  entryComponents: [ProjectSelectModalComponent],
  imports: [SharedModuleModule, ProjectRoutingModule]
})
export class ProjectModule {}
