import { Injectable } from '@angular/core';
import { environment } from '@env/environment';
import { HttpClient } from '@angular/common/http';
import { AuthRequest, AuthResponse } from '@shared/models';
import { Observable } from 'rxjs';

@Injectable()
export class AuthService {
  private baseUrl = environment.baseUrl;

  constructor(private http: HttpClient) {}

  login(req: AuthRequest): Observable<AuthResponse> {
    return this.http.post<AuthResponse>(`${this.baseUrl}/login`, req);
  }

  refresh(): Observable<AuthResponse> {
    return this.http.get<AuthResponse>(`${this.baseUrl}/refresh`);
  }
}
