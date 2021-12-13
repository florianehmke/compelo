import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { CoreModule } from '@core/core.module';
import { SharedModule } from '@shared/shared.module';

import { AppFooterComponent } from './app-footer.component';
import { AppHeaderComponent } from './app-header.component';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ProjectListViewModule } from './project-list-view/project-list-view.module';
import { ProjectViewModule } from './project-view/project-view.module';

@NgModule({
  declarations: [AppComponent, AppHeaderComponent, AppFooterComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    CoreModule,
    SharedModule,
    ProjectViewModule,
    ProjectListViewModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
