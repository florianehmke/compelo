import { Attribute, Component, Input } from '@angular/core';
import { IconName } from '@fortawesome/fontawesome-svg-core';

export type ButtonStyle = 'primary' | 'success' | 'danger';

@Component({
  selector: 'app-button',
  template: `
    <button
      class="btn w-100"
      [type]="type"
      [ngClass]="styleClass"
      [disabled]="disabled"
      [class.disabled]="disabled"
    >
      <span class="text-nowrap"
        >{{ label }}
        <app-icon *ngIf="icon" [icon]="icon" class="ml-1"></app-icon>
      </span>
    </button>
  `
})
export class ButtonComponent {
  @Input()
  label: string;

  @Input()
  style = 'primary';

  @Input()
  disabled = false;

  @Input()
  icon: IconName;

  constructor(@Attribute('type') public type: string = 'button') {}

  get styleClass(): string {
    return `btn-${this.style}`;
  }
}
