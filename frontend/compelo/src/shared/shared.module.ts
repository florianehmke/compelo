import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import {
  NgbModalModule,
  NgbPaginationModule,
  NgbTooltipModule,
} from '@ng-bootstrap/ng-bootstrap';

import { ButtonModule } from './button/button.module';
import { IconModule } from './icon/icon.module';

const modules: any = [
  // angular
  CommonModule,
  FormsModule,
  ReactiveFormsModule,
  // ng-bootstrap
  NgbModalModule,
  NgbPaginationModule,
  NgbTooltipModule,
  // custom
  ButtonModule,
  IconModule,
];

@NgModule({
  imports: [...modules],
  exports: [...modules],
  declarations: [],
  providers: [],
})
export class SharedModule {}
