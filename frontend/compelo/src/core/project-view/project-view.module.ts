import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { StoreModule } from '@ngrx/store';
import { EffectsModule } from '@ngrx/effects';
import { HttpClientModule } from '@angular/common/http';

import { FEATURE_KEY, reducer } from './project-view.reducer';
import { ProjectViewEffects } from './project-view.effects';
import { ProjectViewService } from './project-view.service';

@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    HttpClientModule,
    StoreModule.forFeature(FEATURE_KEY, reducer),
    EffectsModule.forFeature([ProjectViewEffects])
  ],
  providers: [ProjectViewService]
})
export class ProjectViewModule {}
