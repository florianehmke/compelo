import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { CreateProjectRequest, Project } from '@api';
import { Observable } from 'rxjs';

import { environment } from '@env/environment';


@Injectable()
export class ProjectListService {
  private baseUrl = environment.baseUrl + '/projects';

  constructor(private http: HttpClient) {}

  getProjects(): Observable<Project[]> {
    return this.http.get<Project[]>(`${this.baseUrl}`);
  }

  createProject(project: CreateProjectRequest): Observable<Project> {
    return this.http.post<Project>(`${this.baseUrl}`, project);
  }
}
