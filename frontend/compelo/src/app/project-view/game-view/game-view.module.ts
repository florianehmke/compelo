import { NgModule } from '@angular/core';
import { SharedModuleModule } from '@shared/shared.module';

import { GameViewComponent } from './game-view.component';
import { GameViewRoutingModule } from './game-view-routing.module';
import { components } from './components';
import { MatchFormService } from './services/match-form.service';

@NgModule({
  declarations: [GameViewComponent, ...components],
  exports: [],
  imports: [SharedModuleModule, GameViewRoutingModule],
  providers: [MatchFormService]
})
export class GameViewModule {}
