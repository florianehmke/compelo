import { ActivatedRouteSnapshot } from '@angular/router';
import { Action, Store } from '@ngrx/store';
import { Observable, of } from 'rxjs';

import { gameGuidParam } from '@shared/route-params';

import { loadCompetitions } from '.';
import {
  loadGames,
  loadGameStats,
  loadMatches,
  loadPlayers,
  loadPlayerStats,
} from './project.actions';
import { State } from './project.reducer';

export function triggerLoadPlayers(): Observable<Action> {
  return of(loadPlayers());
}

export function triggerLoadGames(): Observable<Action> {
  return of(loadGames());
}

export function triggerLoadMatches(
  store: Store<State>,
  route: ActivatedRouteSnapshot
): Observable<Action> {
  const gameGuid = route.params[gameGuidParam];
  const payload = { payload: { gameGuid: gameGuid } };
  return of(loadMatches(payload));
}

export function triggerLoadPlayerStats(
  store: Store<State>,
  route: ActivatedRouteSnapshot
): Observable<Action> {
  const gameGuid = route.params[gameGuidParam];
  const payload = { payload: { gameGuid: gameGuid } };
  return of(loadPlayerStats(payload));
}

export function triggerLoadGameStats(
  store: Store<State>,
  route: ActivatedRouteSnapshot
): Observable<Action> {
  const gameGuid = route.params[gameGuidParam];
  const payload = { payload: { gameGuid: gameGuid } };
  return of(loadGameStats(payload));
}

export function triggerLoadCompetitions(
  store: Store<State>,
  route: ActivatedRouteSnapshot
): Observable<Action> {
  const gameGuid = route.params[gameGuidParam];
  const payload = { payload: { gameGuid: gameGuid } };
  return of(loadCompetitions(payload));
}
