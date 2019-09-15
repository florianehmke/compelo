import { NgModule } from '@angular/core';
import { SharedModuleModule } from '@shared/shared.module';

import { GameViewComponent } from './game-view.component';
import { GameViewRoutingModule } from './game-view-routing.module';
import { MatchFormService } from './services/match-form.service';
import { components } from './components';
import { pipes } from './pipes';

@NgModule({
  declarations: [GameViewComponent, ...components, ...pipes],
  exports: [],
  imports: [SharedModuleModule, GameViewRoutingModule],
  providers: [MatchFormService]
})
export class GameViewModule {}
