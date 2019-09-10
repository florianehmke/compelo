import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable } from 'rxjs';
import {
  CreateProjectPayload,
  Project,
  SelectProjectPayload,
  TokenPayload
} from '../../shared/models';

@Injectable()
export class ProjectService {
  private baseUrl = environment.baseUrl;

  constructor(private http: HttpClient) {}

  getProjects(): Observable<Project[]> {
    return this.http.get<Project[]>(`${this.baseUrl}/projects`);
  }

  createProject(project: CreateProjectPayload): Observable<Project> {
    return this.http.post<Project>(`${this.baseUrl}/create-project`, project);
  }

  selectProject(project: SelectProjectPayload): Observable<TokenPayload> {
    return this.http.post<TokenPayload>(
      `${this.baseUrl}/select-project`,
      project
    );
  }
}
