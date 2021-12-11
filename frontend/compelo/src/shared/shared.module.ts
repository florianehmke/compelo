import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import {
  NgbModalModule,
  NgbPaginationModule,
  NgbTooltipModule,
} from '@ng-bootstrap/ng-bootstrap';
import { ChartsModule } from 'ng2-charts';

import { ListGroupModule } from './list-group/list-group.module';
import { IconModule } from './icon/icon.module';
import { ButtonModule } from './button/button.module';
import { ToastModule } from './toast/toast.module';
import { StatsBarModule } from './stats-bar/stats-bar.module';

const modules: any = [
  // angular
  CommonModule,
  FormsModule,
  ReactiveFormsModule,

  // ng-bootstrap
  NgbModalModule,
  NgbPaginationModule,
  NgbTooltipModule,

  // chart.js (ng2-charts)
  ChartsModule,

  // custom
  ButtonModule,
  IconModule,
  ListGroupModule,
  ToastModule,
  StatsBarModule,
];

@NgModule({
  imports: [...modules],
  exports: [...modules],
  declarations: [],
  providers: [],
})
export class SharedModule {}
