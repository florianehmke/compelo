import { NgModule } from '@angular/core';
import { SharedModuleModule } from '@shared/shared.module';

import { GameViewComponent } from './game-view.component';
import { GameViewRoutingModule } from './game-view-routing.module';

@NgModule({
  declarations: [GameViewComponent],
  exports: [],
  imports: [SharedModuleModule, GameViewRoutingModule]
})
export class GameViewModule {}
