import { NgModule } from '@angular/core';
import { SharedModuleModule } from '@shared/shared.module';

import { ProjectListViewComponent } from './project-list-view.component';
import { ProjectListViewRoutingModule } from './project-list-view-routing.module';
import { components, entryComponents } from './components';

@NgModule({
  declarations: [ProjectListViewComponent, ...components],
  exports: [],
  entryComponents: [...entryComponents],
  imports: [SharedModuleModule, ProjectListViewRoutingModule]
})
export class ProjectListViewModule {}
