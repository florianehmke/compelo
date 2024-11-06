import { NgModule } from '@angular/core';

import { SharedModule } from '@shared/shared.module';

import { components, entryComponents } from './components';
import { ProjectListViewRoutingModule } from './project-list-view-routing.module';
import { ProjectListViewComponent } from './project-list-view.component';

@NgModule({
  declarations: [ProjectListViewComponent, ...components],
  exports: [],
  imports: [SharedModule, ProjectListViewRoutingModule],
})
export class ProjectListViewModule {}
