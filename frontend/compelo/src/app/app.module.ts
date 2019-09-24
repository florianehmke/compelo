import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { CoreModule } from '@core/core.module';
import { SharedModule } from '@shared/shared.module';
import { ToastModule } from '@shared/toast/toast.module';

import { ProjectViewModule } from './project-view/project-view.module';
import { ProjectListViewModule } from './project-list-view/project-list-view.module';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { AppHeaderComponent } from './app-header.component';

@NgModule({
  declarations: [AppComponent, AppHeaderComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    CoreModule,
    SharedModule,
    ToastModule,
    ProjectViewModule,
    ProjectListViewModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {}
