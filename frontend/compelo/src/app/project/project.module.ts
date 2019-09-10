import { NgModule } from '@angular/core';

import { ProjectListComponent } from './project-list.component';
import { ProjectSelectModalComponent } from './project-select-modal.component';
import { SharedModuleModule } from '../../shared/shared.module';

@NgModule({
  declarations: [ProjectListComponent, ProjectSelectModalComponent],
  exports: [ProjectListComponent, ProjectSelectModalComponent],
  entryComponents: [ProjectSelectModalComponent],
  imports: [SharedModuleModule]
})
export class ProjectModule {}
