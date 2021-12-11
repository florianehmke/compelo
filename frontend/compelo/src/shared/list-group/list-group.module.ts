import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ListGroupComponent } from './list-group.component';
import { ListGroupButtonComponent } from './list-group-button.component';
import { ListGroupItemComponent } from './list-group-item.component';

@NgModule({
  declarations: [
    ListGroupComponent,
    ListGroupButtonComponent,
    ListGroupItemComponent,
  ],
  exports: [
    ListGroupComponent,
    ListGroupButtonComponent,
    ListGroupItemComponent,
  ],
  imports: [CommonModule],
})
export class ListGroupModule {}
