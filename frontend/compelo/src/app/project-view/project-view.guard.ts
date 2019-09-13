import { Injectable } from '@angular/core';
import {
  ActivatedRouteSnapshot,
  CanActivate,
  Router,
  RouterStateSnapshot
} from '@angular/router';
import { tokenForProjectIdExists } from '@shared/jwt';

@Injectable()
export class ProjectViewGuard implements CanActivate {
  constructor(private router: Router) {}

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
    if (tokenForProjectIdExists(parseInt(route.paramMap.get('projectId'), 10))) {
      return true;
    }
    console.warn('token does not belong to project');
    return this.router.parseUrl('/project-list');
  }
}
