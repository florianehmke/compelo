import { ChangeDetectionStrategy, Component, Input } from '@angular/core';
import { IconProp } from '@fortawesome/fontawesome-svg-core';

@Component({
  selector: 'app-icon',
  template: `
    <fa-icon [icon]="icon" [class.icon-btn]="button"></fa-icon>
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
  icon: IconProp;

  @Input()
  button = false;
}
