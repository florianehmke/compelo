import { enableProdMode, importProvidersFrom } from '@angular/core';
import { BrowserModule, bootstrapApplication } from '@angular/platform-browser';
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';


import { CoreModule } from '@core/core.module';
import { SharedModule } from '@shared/shared.module';

import { AppRoutingModule } from './app/app-routing.module';
import { AppComponent } from './app/app.component';
import { ProjectListViewModule } from './app/project-list-view/project-list-view.module';
import { ProjectViewModule } from './app/project-view/project-view.module';
import { environment } from './environments/environment';

if (environment.production) {
  enableProdMode();
}

bootstrapApplication(AppComponent, {
  providers: [
    importProvidersFrom(
      BrowserModule,
      AppRoutingModule,
      CoreModule,
      SharedModule,
      ProjectViewModule,
      ProjectListViewModule
    ),
  ],
}).catch((err) => console.error(err));
