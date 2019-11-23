import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '@env/environment';
import { CreateProjectRequest, Project } from '@api';

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
