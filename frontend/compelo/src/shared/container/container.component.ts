import { Component } from '@angular/core';

@Component({
  selector: 'app-container',
  template: `
    <ng-content></ng-content>
  `,
  styles: [
    `
      :host {
        display: block;
        padding: 12px;
        border-width: 1px;
        border-style: solid;
      }
    `
  ]
})
export class ContainerComponent {}
