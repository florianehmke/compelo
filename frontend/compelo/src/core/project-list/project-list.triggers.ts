import { Observable, of } from 'rxjs';
import { Action } from '@ngrx/store';

import { loadProjects } from './project-list.actions';

export function triggerLoadProjects(): Observable<Action> {
  return of(loadProjects());
}
