import { ChangeDetectionStrategy, Component, Input } from '@angular/core';
import { IconName, IconPrefix } from '@fortawesome/fontawesome-svg-core';

@Component({
  selector: 'app-icon',
  template: `
    <fa-icon [icon]="[prefix, icon]" [class.icon-btn]="button"></fa-icon>
  `,
  changeDetection: ChangeDetectionStrategy.OnPush,
  styles: [
    `
      .icon-btn {
        cursor: pointer;
      }
    `
  ]
})
export class IconComponent {
  @Input()
  icon: IconName;

  @Input()
  prefix: IconPrefix = 'fas';

  @Input()
  button = false;
}
