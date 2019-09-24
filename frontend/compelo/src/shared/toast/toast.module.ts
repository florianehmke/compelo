import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NgbToastModule } from '@ng-bootstrap/ng-bootstrap';

import { ToastComponent } from './toast.component';

@NgModule({
  imports: [CommonModule, NgbToastModule],
  exports: [ToastComponent],
  declarations: [ToastComponent],
  providers: []
})
export class ToastModule {}
