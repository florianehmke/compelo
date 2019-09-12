import { ProjectCardComponent } from './project-card.component';
import { ProjectCreateComponent } from './project-create.component';
import { ProjectGridComponent } from './project-grid.component';
import { ProjectSelectModalComponent } from './project-select-modal.component';

export * from './project-card.component';
export * from './project-create.component';
export * from './project-grid.component';
export * from './project-select-modal.component';

export const entryComponents: any[] = [ProjectSelectModalComponent];

export const components: any[] = [
  ProjectCardComponent,
  ProjectCreateComponent,
  ProjectGridComponent,
  ...entryComponents
];
