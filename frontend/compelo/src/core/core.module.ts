import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActionReducer, StoreModule } from '@ngrx/store';
import { EffectsModule } from '@ngrx/effects';
import { storeLogger } from 'ngrx-store-logger';
import { environment } from '@env/environment';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { routerReducer, StoreRouterConnectingModule } from '@ngrx/router-store';

import { ProjectListModule } from './project-list/project-list.module';
import { ProjectModule } from './project/project.module';
import { AuthInterceptor } from './auth.interceptor';
import { AuthService } from './auth.service';

export function logger(reducer: ActionReducer<any>): any {
  return storeLogger()(reducer);
}

export const metaReducers = environment.production ? [] : [logger];

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
