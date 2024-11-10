import { enableProdMode, importProvidersFrom } from '@angular/core';
import { BrowserModule, bootstrapApplication } from '@angular/platform-browser';
import { provideCharts, withDefaultRegisterables } from 'ng2-charts';

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
    provideCharts(withDefaultRegisterables()),
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
