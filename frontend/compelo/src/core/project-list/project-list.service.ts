import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '@env/environment';
import {
  CreateProjectPayload,
  Project,
  SelectProjectPayload,
  TokenPayload
} from '@shared/models';

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

  selectProject(project: SelectProjectPayload): Observable<TokenPayload> {
    return this.http.post<TokenPayload>(
      `${this.baseUrl}/select-project`,
      project
    );
  }
}
