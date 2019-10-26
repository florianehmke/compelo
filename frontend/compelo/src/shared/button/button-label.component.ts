import { Component, Input } from '@angular/core';
import { IconName } from '@fortawesome/fontawesome-common-types';

@Component({
  selector: 'app-button-label',
  template: `
    <span class="text-nowrap">
      <ng-content></ng-content>
      <app-icon *ngIf="icon" [icon]="icon" class="ml-1"></app-icon>
    </span>
  `
})
export class ButtonLabelComponent {
  @Input()
  icon: IconName;
}
