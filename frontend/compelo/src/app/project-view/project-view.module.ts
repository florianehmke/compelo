import { NgModule } from '@angular/core';
import { LoadingSpinnerModule } from '@shared/loading-spinner';
import { SharedModule } from '@shared/shared.module';
import { components } from './components';
import { ProjectViewRoutingModule } from './project-view-routing.module';
import { ProjectViewComponent } from './project-view.component';
import { ProjectViewGuard } from './project-view.guard';

@NgModule({
  declarations: [ProjectViewComponent, ...components],
  exports: [],
  providers: [ProjectViewGuard],
  imports: [SharedModule, ProjectViewRoutingModule, LoadingSpinnerModule],
})
export class ProjectViewModule {}
