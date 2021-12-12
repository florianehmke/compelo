import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { NgbToastModule } from '@ng-bootstrap/ng-bootstrap';

import { ToastComponent } from './toast.component';

@NgModule({
  imports: [CommonModule, NgbToastModule],
  exports: [ToastComponent],
  declarations: [ToastComponent],
  providers: [],
})
export class ToastModule {}
