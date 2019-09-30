import { Injectable } from '@angular/core';
import { environment } from '@env/environment';
import { HttpClient } from '@angular/common/http';
import { AuthRequest, TokenPayload } from '@shared/models';
import { Observable } from 'rxjs';

@Injectable()
export class AuthService {
  private baseUrl = environment.baseUrl;

  constructor(private http: HttpClient) {}

  login(req: AuthRequest): Observable<TokenPayload> {
    return this.http.post<TokenPayload>(`${this.baseUrl}/login`, req);
  }

  refresh(): Observable<TokenPayload> {
    return this.http.get<TokenPayload>(`${this.baseUrl}/refresh`);
  }
}
