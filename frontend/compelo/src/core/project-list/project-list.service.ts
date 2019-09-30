import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '@env/environment';
import { Project } from '@shared/models';
import { CreateProjectPayload } from '@core/project-list';

@Injectable()
export class ProjectListService {
  private baseUrl = environment.baseUrl + '/projects';

  constructor(private http: HttpClient) {}

  getProjects(): Observable<Project[]> {
    return this.http.get<Project[]>(`${this.baseUrl}`);
  }

  createProject(project: CreateProjectPayload): Observable<Project> {
    return this.http.post<Project>(`${this.baseUrl}`, project);
  }
}
