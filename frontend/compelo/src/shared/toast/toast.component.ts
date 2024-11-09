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
      (hidden)="toastService.remove(toast)"
    >
      {{ toast.message }}
    </ngb-toast>
  `,
  standalone: true,
  imports: [NgFor, NgbToast],
  styles: [
    `
      :host {
        position: fixed;
        top: 0;
        right: 0;
        margin: 0.5em;
        z-index: 1200;
      }
    `,
  ],
})
export class ToastComponent {
  @HostBinding('class.ngb-toasts')
  classNgbToasts = true;

  constructor(public toastService: ToastService) {}
}
