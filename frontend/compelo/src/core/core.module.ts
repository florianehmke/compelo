import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActionReducer, StoreModule } from '@ngrx/store';
import { EffectsModule } from '@ngrx/effects';
import { storeLogger } from 'ngrx-store-logger';

import { environment } from '../environments/environment';
import { ProjectListModule } from './project-list/project-list.module';
import { ProjectViewModule } from './project-view/project-view.module';

export function logger(reducer: ActionReducer<any>): any {
  return storeLogger()(reducer);
}

export const metaReducers = environment.production ? [] : [logger];

@NgModule({
  imports: [
    StoreModule.forRoot({}, { metaReducers }),
    EffectsModule.forRoot([]),
    CommonModule,
    ProjectListModule,
    ProjectViewModule
  ]
})
export class CoreModule {}
