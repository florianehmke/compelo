import { NgModule } from '@angular/core';
import { SharedModule } from '@shared/shared.module';

import { GameViewComponent } from './game-view.component';
import { GameViewRoutingModule } from './game-view-routing.module';
import { MatchFormService } from './services/match-form.service';
import { components } from './components';
import { pipes } from './pipes';

@NgModule({
  declarations: [GameViewComponent, ...components, ...pipes],
  exports: [],
  imports: [SharedModule, GameViewRoutingModule],
  providers: [MatchFormService],
})
export class GameViewModule {}
