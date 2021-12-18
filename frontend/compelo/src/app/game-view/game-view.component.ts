import { Component } from '@angular/core';
import { getLoadedByActionTypeOf } from '@core/app';
import {
  createMatch,
  filterMatches,
  getGameStats,
  getMatches,
  getPlayers,
  getPlayerStats,
  loadGameStats,
  loadMatches,
  loadPlayerStats,
  State,
} from '@core/project';
import { CreateMatchRequest } from '@generated/api';
import { Store } from '@ngrx/store';
import { combineLatest } from 'rxjs';
import { map } from 'rxjs/operators';
import {
  MatchFormService,
  MatchFormSettings,
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
      <div class="col-md-6">
        <app-match-list
          [matches]="matches$ | async"
          [isLoaded]="matchesLoaded$ | async"
          (filterChange)="onFilterChange($event)"
        ></app-match-list>
      </div>
      <div class="col-md-6">
        <app-stats
          [isLoaded]="statsLoaded$ | async"
          [players]="playerStats$ | async"
          [gameStats]="gameStats$ | async"
        ></app-stats>
      </div>
    </div>
  `,
})
export class GameViewComponent {
  players$ = this.store.select(getPlayers);

  matches$ = this.store.select(getMatches);

  matchesLoaded$ = this.store.select(getLoadedByActionTypeOf(loadMatches));

  playerStats$ = this.store.select(getPlayerStats);

  gameStats$ = this.store.select(getGameStats);

  statsLoaded$ = combineLatest([
    this.store.select(getLoadedByActionTypeOf(loadPlayerStats)),
    this.store.select(getLoadedByActionTypeOf(loadGameStats)),
  ]).pipe(
    map(
      ([loadedPlayerStats, loadedGameStats]) =>
        loadedPlayerStats && loadedGameStats
    )
  );

  formGroup = this.formService.buildForm({ teamSize: 1, teamCount: 2 });
  showSettings = false;

  constructor(
    private store: Store<State>,
    private formService: MatchFormService
  ) {}

  onFormSettingsChanged(settings: MatchFormSettings) {
    this.formGroup = this.formService.buildForm(settings);
  }

  onMatchCreated(payload: CreateMatchRequest) {
    this.store.dispatch(createMatch({ payload }));
  }

  onFilterChange(filter: string) {
    this.store.dispatch(filterMatches({ payload: { filter } }));
  }
}
