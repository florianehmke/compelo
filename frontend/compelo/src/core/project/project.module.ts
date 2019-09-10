import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ProjectService } from './project.service';
import { StoreModule } from '@ngrx/store';
import { FEATURE_KEY, reducer } from './project.reducer';

@NgModule({
  declarations: [],
  imports: [CommonModule, StoreModule.forFeature(FEATURE_KEY, reducer)],
  providers: [ProjectService]
})
export class ProjectModule {}
