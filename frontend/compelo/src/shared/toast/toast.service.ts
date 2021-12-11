import { Injectable } from '@angular/core';

@Injectable({ providedIn: 'root' })
export class ToastService {
  toasts: any[] = [];

  info(message: string) {
    this.show(message);
  }

  success(message: string) {
    this.show(message, { classname: 'bg-success text-light' });
  }

  danger(message: string) {
    this.show(message, { classname: 'bg-danger text-light' });
  }

  remove(toast) {
    this.toasts = this.toasts.filter((t) => t !== toast);
  }

  private show(message: string, options: any = {}) {
    this.toasts.push({ message, ...options });
  }
}
