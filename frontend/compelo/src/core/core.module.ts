import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { StoreModule } from '@ngrx/store';

import { ProjectModule } from './project/project.module';

@NgModule({
  imports: [CommonModule, StoreModule.forRoot({}), ProjectModule]
})
export class CoreModule {}
