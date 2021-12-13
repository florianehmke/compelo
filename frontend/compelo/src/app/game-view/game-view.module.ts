import { NgModule } from '@angular/core';

import { SharedModule } from '@shared/shared.module';

import { components } from './components';
import { GameViewRoutingModule } from './game-view-routing.module';
import { GameViewComponent } from './game-view.component';
import { pipes } from './pipes';
import { MatchFormService } from './services/match-form.service';

@NgModule({
  declarations: [GameViewComponent, ...components, ...pipes],
  exports: [],
  imports: [SharedModule, GameViewRoutingModule],
  providers: [MatchFormService],
})
export class GameViewModule {}
