import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  template: `
    <div class="header lead">> compelo</div>
    <div class="container">
      <router-outlet></router-outlet>
    </div>
  `,

  styles: [`
    .header {
      height: 48px;
      padding: 8px;
      background-color: cornflowerblue;
      color: white;
      margin-bottom: 32px;
    }
  `]
})
export class AppComponent {}
