import { Injectable } from '@angular/core';
import {
  ActivatedRouteSnapshot,
  CanActivate,
  Router,
  RouterStateSnapshot,
} from '@angular/router';

import { tokenForProjectIdExists } from '@shared/jwt';
import { projectGuidParam } from '@shared/route-params';

@Injectable()
export class ProjectViewGuard implements CanActivate {
  constructor(private router: Router) {}

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
    const projectGuid = route.paramMap.get(projectGuidParam);
    if (tokenForProjectIdExists(projectGuid)) {
      return true;
    }
    console.warn('token does not belong to project');
    return this.router.parseUrl('/project-list');
  }
}
