import { Component } from '@angular/core';

@Component({
  selector: 'app-list-group-button',
  template: `
    <button type="button" class="list-group-item list-group-item-action">
      <ng-content></ng-content>
    </button>
  `,
})
export class ListGroupButtonComponent {}
