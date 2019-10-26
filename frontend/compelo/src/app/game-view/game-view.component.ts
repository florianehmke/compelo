import { Component } from '@angular/core';
import {
  createMatch,
  CreateMatchPayload,
  filterMatches,
  getMatches,
  getPlayers,
  getPlayerStats,
  State
} from '@core/project';
import { Store } from '@ngrx/store';

import {
  MatchFormService,
  MatchFormSettings
} from './services/match-form.service';

@Component({
  template: `
    <div class="row">
      <div class="col-md">
        <app-match-create
          [formGroup]="formGroup"
          [players]="players$ | async"
          (matchCreated)="onMatchCreated($event)"
          (settingsClick)="showSettings = !showSettings"
        ></app-match-create>
      </div>
      <div class="col-md" *ngIf="showSettings">
        <app-match-settings
          (settingsClose)="showSettings = !showSettings"
          (formSettingsChanged)="onFormSettingsChanged($event)"
        ></app-match-settings>
      </div>
    </div>
    <hr />
    <div class="row">
      <div class="col-md-6" *ngIf="matches$ | async as matches">
        <app-match-list
          [matches]="matches"
          (filterChange)="onFilterChange($event)"
        ></app-match-list>
      </div>
      <div class="col-md-6" *ngIf="playerStats$ | async as players">
        <app-player-stats [players]="players"></app-player-stats>
      </div>
    </div>
  `
})
export class GameViewComponent {
  players$ = this.store.select(getPlayers);
  playerStats$ = this.store.select(getPlayerStats);
  matches$ = this.store.select(getMatches);

  formGroup = this.formService.buildForm({ teamSize: 1, teamCount: 2 });
  showSettings = false;

  constructor(
    private store: Store<State>,
    private formService: MatchFormService
  ) {}

  onFormSettingsChanged(settings: MatchFormSettings) {
    this.formGroup = this.formService.buildForm(settings);
  }

  onMatchCreated(payload: CreateMatchPayload) {
    this.store.dispatch(createMatch({ payload }));
  }

  onFilterChange(filter: string) {
    this.store.dispatch(filterMatches({ payload: { filter } }));
  }
}
