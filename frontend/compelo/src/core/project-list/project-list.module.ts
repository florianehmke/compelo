import { CommonModule } from '@angular/common';
import {
  provideHttpClient,
  withInterceptorsFromDi,
} from '@angular/common/http';
import { NgModule } from '@angular/core';
import { EffectsModule } from '@ngrx/effects';
import { StoreModule } from '@ngrx/store';

import { ProjectListEffects } from './project-list.effects';
import { FEATURE_KEY, reducer } from './project-list.reducer';
import { ProjectListService } from './project-list.service';

@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    StoreModule.forFeature(FEATURE_KEY, reducer),
    EffectsModule.forFeature([ProjectListEffects]),
  ],
  providers: [ProjectListService, provideHttpClient(withInterceptorsFromDi())],
})
export class ProjectListModule {}
