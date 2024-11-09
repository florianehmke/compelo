import { Component } from '@angular/core';

@Component({
  selector: 'app-list-group-button',
  template: `
    <button type="button" class="bg-white list-group-item list-group-item-action">
      <ng-content></ng-content>
    </button>
  `,
  standalone: true,
})
export class ListGroupButtonComponent {}
