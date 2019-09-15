import { Component } from '@angular/core';
import {
  createMatch,
  CreateMatchPayload,
  getMatches,
  getPlayers,
  State
} from '@core/project';
import { Store } from '@ngrx/store';
import {
  MatchFormService,
  MatchFormSettings
} from './services/match-form.service';

@Component({
  selector: 'app-game-view',
  template: `
    <p class="lead">
      Change Match Settings
    </p>
    <app-match-settings
      (formSettingsChanged)="onFormSettingsChanged($event)"
    ></app-match-settings>
    <hr />
    <p class="lead">
      Create Match
    </p>
    <app-match-create
      [formGroup]="formGroup"
      [players]="players$ | async"
      (matchCreated)="onMatchCreated($event)"
    ></app-match-create>
    <hr />
    <p class="lead">
      Recent Matches
    </p>
    <app-match-list [matches]="matches$ | async"> </app-match-list>
  `
})
export class GameViewComponent {
  players$ = this.store.select(getPlayers);
  matches$ = this.store.select(getMatches);

  formGroup = this.formService.buildForm({ teamSize: 1, teamCount: 2 });

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
}
