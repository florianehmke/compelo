import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  template: `
    <div class="header">> compelo</div>
    <div class="container">
      <router-outlet></router-outlet>
    </div>
  `,
  styles: []
})
export class AppComponent {}
