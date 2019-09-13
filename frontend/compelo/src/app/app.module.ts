import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { CoreModule } from '@core/core.module';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ProjectViewModule } from './project-view/project-view.module';
import { ProjectListViewModule } from './project-list-view/project-list-view.module';

@NgModule({
  declarations: [AppComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    CoreModule,
    ProjectViewModule,
    ProjectListViewModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {}
