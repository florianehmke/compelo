import { Observable, of } from 'rxjs';
import { Action } from '@ngrx/store';

import { loadGames, loadMatches, loadPlayers } from './project.actions';

export function triggerLoadPlayers(): Observable<Action> {
  return of(loadPlayers());
}

export function triggerLoadGames(): Observable<Action> {
  return of(loadGames({ payload: { loadMatches: false } }));
}

export function triggerLoadGamesAndMatches(): Observable<Action> {
  return of(loadGames({ payload: { loadMatches: true } }));
}
