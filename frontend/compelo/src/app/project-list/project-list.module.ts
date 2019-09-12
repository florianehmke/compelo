import { NgModule } from '@angular/core';

import { ProjectListComponent } from './project-list.component';
import { SharedModuleModule } from '../../shared/shared.module';
import { ProjectListRoutingModule } from './project-list-routing.module';
import { components, entryComponents } from './components';

@NgModule({
  declarations: [ProjectListComponent, ...components],
  exports: [],
  entryComponents: [...entryComponents],
  imports: [SharedModuleModule, ProjectListRoutingModule]
})
export class ProjectListModule {}
