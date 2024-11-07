import { Directive, HostBinding } from '@angular/core';

@Directive({
  selector: '[appPrimary]',
  standalone: true,
})
export class ButtonPrimaryDirective {
  @HostBinding('class.btn') btn = true;
  @HostBinding('class.btn-primary') btnPrimary = true;
}
