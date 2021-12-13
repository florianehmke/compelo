import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { IconModule } from '@shared/icon/icon.module';

import { ButtonLabelComponent } from './button-label.component';
import { ButtonPrimaryDirective } from './button-primary.directive';

@NgModule({
  imports: [CommonModule, IconModule],
  exports: [ButtonLabelComponent, ButtonPrimaryDirective],
  declarations: [ButtonLabelComponent, ButtonPrimaryDirective],
  providers: [],
})
export class ButtonModule {}
