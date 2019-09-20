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

const modules: any = [
  CommonModule,
  FormsModule,
  ReactiveFormsModule,
  NgbModalModule,
  NgbPaginationModule,
  ButtonModule,
  IconModule,
  ListGroupModule
];

@NgModule({
  imports: [...modules],
  exports: [...modules],
  declarations: [],
  providers: []
})
export class SharedModuleModule {}
