import { Component } from '@angular/core';

@Component({
  selector: 'app-list-group',
  template: `
    <div class="list-group">
      <ng-content></ng-content>
    </div>
  `,
})
export class ListGroupComponent {}
