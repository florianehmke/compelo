import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
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
    HttpClientModule,
    StoreModule.forFeature(FEATURE_KEY, reducer),
    EffectsModule.forFeature([ProjectEffects]),
  ],
  providers: [ProjectService],
})
export class ProjectModule {}
