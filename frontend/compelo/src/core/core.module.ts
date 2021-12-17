import { CommonModule } from '@angular/common';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { environment } from '@env/environment';
import { EffectsModule } from '@ngrx/effects';
import { routerReducer, StoreRouterConnectingModule } from '@ngrx/router-store';
import { ActionReducer, StoreModule } from '@ngrx/store';
import { storeLogger } from 'ngrx-store-logger';
import { AppEffects, appFeatureKey, reducer as appReducer } from './app';
import { AuthInterceptor } from './auth.interceptor';
import { AuthService } from './auth.service';
import { ProjectListModule } from './project-list/project-list.module';
import { ProjectModule } from './project/project.module';

export function logger(reducer: ActionReducer<any>): any {
  return storeLogger()(reducer);
}

export const metaReducers = environment.production ? [] : [logger];

@NgModule({
  imports: [
    CommonModule,
    StoreModule.forRoot(
      { router: routerReducer, [appFeatureKey]: appReducer },
      { metaReducers }
    ),
    StoreRouterConnectingModule.forRoot(),
    EffectsModule.forRoot([AppEffects]),
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
