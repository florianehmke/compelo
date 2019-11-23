import { Injectable } from '@angular/core';
import { environment } from '@env/environment';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { AuthRequest, AuthResponse } from '@api';

@Injectable()
export class AuthService {
  private baseUrl = environment.baseUrl;

  constructor(private http: HttpClient) {}

  login(req: AuthRequest): Observable<AuthResponse> {
    return this.http.post<AuthResponse>(`${this.baseUrl}/login`, req);
  }

  refresh(): Observable<AuthResponse> {
    return this.http.post<AuthResponse>(`${this.baseUrl}/refresh`, {});
  }
}
