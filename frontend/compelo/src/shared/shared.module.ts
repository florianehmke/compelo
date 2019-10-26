import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import {
  NgbModalModule,
  NgbPaginationModule
} from '@ng-bootstrap/ng-bootstrap';

import { ListGroupModule } from './list-group/list-group.module';
import { IconModule } from './icon/icon.module';
import { ButtonModule } from './button/button.module';
import { ToastModule } from './toast/toast.module';
import { ChartsModule } from 'ng2-charts';

const modules: any = [
  // angular
  CommonModule,
  FormsModule,
  ReactiveFormsModule,

  // ng-bootstrap
  NgbModalModule,
  NgbPaginationModule,

  // chart.js (ng2-charts)
  ChartsModule,

  // custom
  ButtonModule,
  IconModule,
  ListGroupModule,
  ToastModule
];

@NgModule({
  imports: [...modules],
  exports: [...modules],
  declarations: [],
  providers: []
})
export class SharedModule {}
