import { Directive, HostBinding } from '@angular/core';

@Directive({ selector: '[appPrimary]' })
export class ButtonPrimaryDirective {
  @HostBinding('class.btn') btn = true;
  @HostBinding('class.btn-primary') btnPrimary = true;
}
