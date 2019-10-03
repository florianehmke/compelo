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
    const projectId = route.paramMap.get('projectId');
    if (tokenForProjectIdExists(projectId)) {
      return true;
    }
    console.warn('token does not belong to project');
    return this.router.parseUrl('/project-list');
  }
}
