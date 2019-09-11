import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CoreModule } from '../core/core.module';
import { ProjectViewModule } from './project-view/project-view.module';
import { ProjectListModule } from './project-list/project-list.module';

@NgModule({
  declarations: [AppComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    CoreModule,
    ProjectViewModule,
    ProjectListModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {}
