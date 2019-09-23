import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '@env/environment';
import { AuthRequest, Project, TokenPayload } from '@shared/models';
import { CreateProjectPayload } from '@core/project-list';

@Injectable()
export class ProjectListService {
  private baseUrl = environment.baseUrl;

  constructor(private http: HttpClient) {}

  getProjects(): Observable<Project[]> {
    return this.http.get<Project[]>(`${this.baseUrl}/projects`);
  }

  createProject(project: CreateProjectPayload): Observable<Project> {
    return this.http.post<Project>(`${this.baseUrl}/create-project`, project);
  }
}
