import { CommonModule } from '@angular/common';
import {
  provideHttpClient,
  withInterceptorsFromDi,
} from '@angular/common/http';
import { NgModule } from '@angular/core';
import { EffectsModule } from '@ngrx/effects';
import { StoreModule } from '@ngrx/store';

import { ProjectEffects } from './project.effects';
import { FEATURE_KEY, reducer } from './project.reducer';
import { ProjectService } from './project.service';

@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    StoreModule.forFeature(FEATURE_KEY, reducer),
    EffectsModule.forFeature([ProjectEffects]),
  ],
  providers: [ProjectService, provideHttpClient(withInterceptorsFromDi())],
})
export class ProjectModule {}
