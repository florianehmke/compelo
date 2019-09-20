import { Observable, of } from 'rxjs';
import { Action, Store } from '@ngrx/store';
import { ActivatedRouteSnapshot } from '@angular/router';
import { gameIdParam } from '@shared/route-params';

import {
  loadGames,
  loadMatches,
  loadPlayers,
  loadPlayersWithStats
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
  const gameId = route.params[gameIdParam];
  const payload = { payload: { gameId: parseInt(gameId, 10) } };
  return of(loadMatches(payload));
}

export function triggerLoadPlayersWithStats(
  store: Store<State>,
  route: ActivatedRouteSnapshot
): Observable<Action> {
  const gameId = route.params[gameIdParam];
  const payload = { payload: { gameId: parseInt(gameId, 10) } };
  return of(loadPlayersWithStats(payload));
}
