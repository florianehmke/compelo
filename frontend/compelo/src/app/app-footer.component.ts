import { Component } from '@angular/core';

import { APP_VERSION } from './version';

@Component({
  selector: 'app-footer',
  template: `<div class="my-3 text-muted">Build: {{ appVersion }}</div>`,
  styles: [
    `
      :host {
        align-self: center;
      }
    `,
  ],
})
export class AppFooterComponent {
  appVersion = APP_VERSION;
}
