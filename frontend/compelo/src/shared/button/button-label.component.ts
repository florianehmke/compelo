import { NgIf } from '@angular/common';
import { Component, Input } from '@angular/core';
import { IconName } from '@fortawesome/fontawesome-common-types';

import { IconComponent } from '../icon/icon.component';

@Component({
  selector: 'app-button-label',
  template: `
    <span class="text-nowrap">
      <ng-content></ng-content>
      <app-icon *ngIf="icon" [icon]="icon" class="ml-1"></app-icon>
    </span>
  `,
  standalone: true,
  imports: [NgIf, IconComponent],
})
export class ButtonLabelComponent {
  @Input()
  icon: IconName;
}
