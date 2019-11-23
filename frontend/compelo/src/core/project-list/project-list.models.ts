import { AuthRequest, Project } from '@api';

export interface SelectProjectPayload {
  request: AuthRequest;
  project: Project;
}
