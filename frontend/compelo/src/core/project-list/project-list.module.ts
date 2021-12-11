import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { StoreModule } from '@ngrx/store';
import { EffectsModule } from '@ngrx/effects';
import { HttpClientModule } from '@angular/common/http';

import { ProjectListService } from './project-list.service';
import { FEATURE_KEY, reducer } from './project-list.reducer';
import { ProjectListEffects } from './project-list.effects';

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
