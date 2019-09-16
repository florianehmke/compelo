import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NgbModalModule } from '@ng-bootstrap/ng-bootstrap';

import { ListGroupModule } from './list-group/list-group.module';

const modules: any = [
  CommonModule,
  FormsModule,
  ReactiveFormsModule,
  NgbModalModule,
  ListGroupModule
];

@NgModule({
  imports: [...modules],
  exports: [...modules],
  declarations: [],
  providers: []
})
export class SharedModuleModule {}
