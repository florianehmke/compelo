import { Injectable } from '@angular/core';
import {
  ActivatedRouteSnapshot,
  Router,
  RouterStateSnapshot,
} from '@angular/router';

import { tokenForProjectIdExists } from '@shared/jwt';

@Injectable()
export class ProjectViewGuard {
  constructor(private router: Router) {}

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
    const projectId = parseInt(route.paramMap.get('projectId'), 10);
    if (tokenForProjectIdExists(projectId)) {
      return true;
    }
    console.warn('token does not belong to project');
    return this.router.parseUrl('/project-list');
  }
}
