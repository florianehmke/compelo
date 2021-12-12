import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { CoreModule } from '@core/core.module';
import { SharedModule } from '@shared/shared.module';

import { ProjectViewModule } from './project-view/project-view.module';
import { ProjectListViewModule } from './project-list-view/project-list-view.module';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { AppHeaderComponent } from './app-header.component';
import { AppFooterComponent } from './app-footer.component';

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
