import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Resolve } from '@angular/router';
import { Action, Store } from '@ngrx/store';
import { Observable, of } from 'rxjs';
import { map } from 'rxjs/operators';
import { isArray } from 'util';

export type Actions = Action | Action[];

export type RouteActionFactory<T> = (
  store?: Store<T>,
  route?: ActivatedRouteSnapshot
) => Observable<Actions>;

interface RouteData<T> {
  actionFactory: Array<RouteActionFactory<T>>;
}

@Injectable()
export class ActionResolver<T> implements Resolve<boolean> {
  constructor(private store: Store<T>) {}

  resolve(route: ActivatedRouteSnapshot): Observable<boolean> {
    const { actionFactory } = route.routeConfig.data as RouteData<T>;

    actionFactory.forEach((factory: RouteActionFactory<T>) =>
      factory(this.store, route)
        .pipe(map(actions => (isArray(actions) ? actions : [actions])))
        .subscribe((actions: Action[]) => {
          actions.forEach((action: Action) => this.store.dispatch(action));
        })
    );

    return of(true);
  }
}
