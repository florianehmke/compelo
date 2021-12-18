import { NgModule } from '@angular/core';

import { LoadingSpinnerModule } from '@shared/loading-spinner';
import { SharedModule } from '@shared/shared.module';

import { components, entryComponents } from './components';
import { ProjectListViewRoutingModule } from './project-list-view-routing.module';
import { ProjectListViewComponent } from './project-list-view.component';

@NgModule({
  declarations: [ProjectListViewComponent, ...components],
  exports: [],
  entryComponents: [...entryComponents],
  imports: [SharedModule, ProjectListViewRoutingModule, LoadingSpinnerModule],
})
export class ProjectListViewModule {}
