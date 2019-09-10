import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';

@Injectable()
export class ProjectService {
  private baseUrl = environment.baseUrl;

  constructor(private http: HttpClient) {}
}
