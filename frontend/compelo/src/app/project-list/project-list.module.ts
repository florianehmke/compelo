import { NgModule } from '@angular/core';

import { ProjectListComponent } from './project-list.component';
import { ProjectSelectModalComponent } from './project-select-modal.component';
import { SharedModuleModule } from '../../shared/shared.module';
import { ProjectListRoutingModule } from './project-list-routing.module';
import { ProjectGridComponent } from './components/project-grid.component';
import { ProjectCardComponent } from './components/project-card.component';
import { ProjectCreateComponent } from './components/project-create.component';

@NgModule({
  declarations: [
    ProjectListComponent,
    ProjectSelectModalComponent,
    ProjectGridComponent,
    ProjectCardComponent,
    ProjectCreateComponent
  ],
  exports: [],
  entryComponents: [ProjectSelectModalComponent],
  imports: [SharedModuleModule, ProjectListRoutingModule]
})
export class ProjectListModule {}
