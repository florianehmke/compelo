import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { StoreModule } from '@ngrx/store';
import { EffectsModule } from '@ngrx/effects';
import { HttpClientModule } from '@angular/common/http';

import { ProjectService } from './project.service';
import { FEATURE_KEY, reducer } from './project.reducer';
import { ProjectEffects } from './project.effects';

@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    HttpClientModule,
    StoreModule.forFeature(FEATURE_KEY, reducer),
    EffectsModule.forFeature([ProjectEffects])
  ],
  providers: [ProjectService]
})
export class ProjectModule {}
