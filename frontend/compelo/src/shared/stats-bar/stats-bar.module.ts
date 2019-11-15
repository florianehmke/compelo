import { NgModule } from '@angular/core';
import { StatsBarComponent } from '@shared/stats-bar/stats-bar.component';
import { CommonModule } from '@angular/common';

@NgModule({
  imports: [CommonModule],
  exports: [StatsBarComponent],
  declarations: [StatsBarComponent],
  providers: []
})
export class StatsBarModule {}
