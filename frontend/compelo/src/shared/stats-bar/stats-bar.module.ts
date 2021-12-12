import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { StatsBarComponent } from '@shared/stats-bar/stats-bar.component';

@NgModule({
  imports: [CommonModule],
  exports: [StatsBarComponent],
  declarations: [StatsBarComponent],
  providers: [],
})
export class StatsBarModule {}
