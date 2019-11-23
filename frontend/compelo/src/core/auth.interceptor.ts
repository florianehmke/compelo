import {
  HttpErrorResponse,
  HttpHandler,
  HttpInterceptor,
  HttpRequest
} from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, throwError } from 'rxjs';
import { loadToken, removeToken, storeToken } from '@shared/jwt';
import { catchError, filter, switchMap, take } from 'rxjs/operators';
import { AuthService } from '@core/auth.service';
import { Router } from '@angular/router';
import { ToastService } from '@shared/toast';
import { AuthResponse } from '@shared/models';

@Injectable()
export class AuthInterceptor implements HttpInterceptor {
  constructor(
    private authService: AuthService,
    private router: Router,
    private toastService: ToastService
  ) {}

  private isRefreshing = false;
  private refreshTokenSubject = new BehaviorSubject<any>(null);

  intercept(req: HttpRequest<any>, next: HttpHandler) {
    const token = loadToken();
    req = token ? this.addToken(req, token) : req;

    return next.handle(req).pipe(
      catchError(error => {
        if (req.url.includes('refresh')) {
          this.handleFailedRefresh();
        }
        if (error instanceof HttpErrorResponse && error.status === 401) {
          return this.handle401Error(req, next);
        } else {
          return throwError(error);
        }
      })
    );
  }

  private handle401Error(request: HttpRequest<any>, next: HttpHandler) {
    if (!this.isRefreshing) {
      this.isRefreshing = true;
      this.refreshTokenSubject.next(null);

      return this.authService.refresh().pipe(
        switchMap(({ token }: AuthResponse) => {
          this.isRefreshing = false;
          this.refreshTokenSubject.next(token);
          storeToken(token);
          return next.handle(this.addToken(request, token));
        })
      );
    } else {
      return this.refreshTokenSubject.pipe(
        filter(token => token != null),
        take(1),
        switchMap(token => {
          return next.handle(this.addToken(request, token));
        })
      );
    }
  }

  private handleFailedRefresh() {
    removeToken();
    this.toastService.danger('Token expired! Login again..');
    this.router.navigate(['/']);
  }

  private addToken(request: HttpRequest<any>, token: string): HttpRequest<any> {
    return request.clone({
      setHeaders: {
        Authorization: `Bearer ${token}`
      }
    });
  }
}
