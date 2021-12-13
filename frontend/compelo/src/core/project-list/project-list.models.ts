import { AuthRequest, Project } from '@generated/api';

export interface SelectProjectPayload {
  request: AuthRequest;
  project: Project;
}
