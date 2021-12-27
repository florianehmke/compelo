import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { environment } from '@env/environment';
import { CreateProjectRequest, Project, Response } from '@generated/api';

@Injectable()
export class ProjectListService {
  private baseUrl = environment.baseUrl + '/projects';

  constructor(private http: HttpClient) {}

  getProjects(): Observable<Project[]> {
    return this.http.get<Project[]>(`${this.baseUrl}`);
  }

  createProject(project: CreateProjectRequest): Observable<Response> {
    return this.http.post<Response>(`${this.baseUrl}`, project);
  }
}
