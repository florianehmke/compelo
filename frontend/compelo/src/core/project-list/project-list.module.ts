import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
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
    HttpClientModule,
    StoreModule.forFeature(FEATURE_KEY, reducer),
    EffectsModule.forFeature([ProjectListEffects]),
  ],
  providers: [ProjectListService],
})
export class ProjectListModule {}
