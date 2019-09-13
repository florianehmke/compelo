import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  template: `
    <div class="header lead"><a [routerLink]="['/']">> compelo</a></div>
    <div class="container">
      <router-outlet></router-outlet>
    </div>
  `,

  styles: [
    `
      .header {
        height: 48px;
        padding: 8px;
        background-color: cornflowerblue;
        margin-bottom: 32px;
      }
      a {
        color: white;
      }
    `
  ]
})
export class AppComponent {}
