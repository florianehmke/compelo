import { Component } from '@angular/core';

@Component({
  selector: 'app-list-group-item',
  template: `
    <button
      type="button"
      class="disabled list-group-item list-group-item-action"
    >
      <ng-content></ng-content>
    </button>
  `,
})
export class ListGroupItemComponent {}
