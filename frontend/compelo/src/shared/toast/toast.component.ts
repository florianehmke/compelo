import { NgFor } from '@angular/common';
import { Component, HostBinding } from '@angular/core';
import { NgbToast } from '@ng-bootstrap/ng-bootstrap';

import { ToastService } from '@shared/toast/toast.service';

@Component({
  selector: 'app-toast',
  template: `
    <ngb-toast
      *ngFor="let toast of toastService.toasts"
      [class]="toast.classname"
      [autohide]="true"
      [delay]="toast.delay || 5000"
      (hide)="toastService.remove(toast)"
    >
      {{ toast.message }}
    </ngb-toast>
  `,
  standalone: true,
  imports: [NgFor, NgbToast],
})
export class ToastComponent {
  @HostBinding('class.ngb-toasts')
  classNgbToasts = true;

  constructor(public toastService: ToastService) {}
}
