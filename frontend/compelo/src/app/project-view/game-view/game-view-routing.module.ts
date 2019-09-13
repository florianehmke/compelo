import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { GameViewComponent } from './game-view.component';

const routes: Routes = [
  {
    path: '',
    component: GameViewComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class GameViewRoutingModule {}
