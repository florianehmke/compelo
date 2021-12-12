import { Component } from '@angular/core';

import { APP_BUILD_DATE, APP_VERSION } from './version';

@Component({
  selector: 'app-footer',
  template: `<div class="my-3 d-flex flex-column align-items-center">
    <small class="text-muted">Build: {{ appVersion }}</small>
    <small class="text-muted">Date: {{ appBuildDate }}</small>
  </div>`,
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
  appBuildDate = APP_BUILD_DATE;
}
