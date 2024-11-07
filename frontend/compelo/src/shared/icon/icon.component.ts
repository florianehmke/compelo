import { Component, Input } from '@angular/core';
import { FaIconComponent } from '@fortawesome/angular-fontawesome';
import { IconName, IconPrefix } from '@fortawesome/fontawesome-svg-core';

@Component({
  selector: 'app-icon',
  template: `
    <fa-icon [icon]="[prefix, icon]" [class.icon-btn]="button"></fa-icon>
  `,
  styles: [
    `
      .icon-btn {
        cursor: pointer;
      }
    `,
  ],
  standalone: true,
  imports: [FaIconComponent],
})
export class IconComponent {
  @Input()
  icon: IconName;

  @Input()
  prefix: IconPrefix = 'fas';

  @Input()
  button = false;
}
