import { CommonModule } from '@angular/common';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { EffectsModule } from '@ngrx/effects';
import { routerReducer, StoreRouterConnectingModule } from '@ngrx/router-store';
import { StoreModule } from '@ngrx/store';

import { environment } from '@env/environment';

import { AuthInterceptor } from './auth.interceptor';
import { AuthService } from './auth.service';
import { ProjectListModule } from './project-list/project-list.module';
import { ProjectModule } from './project/project.module';

export const metaReducers = environment.production ? [] : [];

@NgModule({
  imports: [
    CommonModule,
    StoreModule.forRoot({ router: routerReducer }, { metaReducers }),
    StoreRouterConnectingModule.forRoot(),
    EffectsModule.forRoot([]),
    ProjectListModule,
    ProjectModule,
  ],
  providers: [
    AuthService,
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi: true,
    },
  ],
})
export class CoreModule {}
