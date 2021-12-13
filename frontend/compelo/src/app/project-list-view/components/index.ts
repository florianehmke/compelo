import { ProjectCreateComponent } from './project-create.component';
import { ProjectListComponent } from './project-list.component';
import { ProjectSelectModalComponent } from './project-select-modal.component';

export * from './project-create.component';
export * from './project-list.component';
export * from './project-select-modal.component';

export const entryComponents: any[] = [ProjectSelectModalComponent];

export const components: any[] = [
  ProjectCreateComponent,
  ProjectListComponent,
  ...entryComponents,
];
