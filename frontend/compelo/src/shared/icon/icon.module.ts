import { NgModule } from '@angular/core';
import {
  FaIconLibrary,
  FontAwesomeModule
} from '@fortawesome/angular-fontawesome';
import {
  faPlus,
  faSave,
  faTimes,
  faWrench
} from '@fortawesome/free-solid-svg-icons';

import { IconComponent } from './icon.component';

@NgModule({
  imports: [FontAwesomeModule],
  exports: [IconComponent],
  declarations: [IconComponent],
  providers: []
})
export class IconModule {
  constructor(library: FaIconLibrary) {
    library.addIcons(faWrench, faSave, faTimes, faPlus);
  }
}
