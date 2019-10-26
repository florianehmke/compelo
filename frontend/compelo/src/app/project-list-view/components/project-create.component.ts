import { Component, EventEmitter, Output } from '@angular/core';
import { CreateProjectPayload } from '@core/project-list';

@Component({
  selector: 'app-project-create',
  template: `
    <p class="lead">Create / Select Project</p>
    <form (ngSubmit)="onSubmit(); form.reset()" #form="ngForm">
      <div class="row">
        <div class="col-md-5">
          <input
            type="text"
            placeholder="Project Name"
            class="form-control"
            name="name"
            required
            [(ngModel)]="name"
          />
        </div>
        <div class="col-md-5">
          <input
            type="password"
            placeholder="Password"
            class="form-control"
            name="password"
            required
            [(ngModel)]="password"
          />
        </div>
        <div class="col-md-2 d-flex justify-content-end">
          <button type="submit" appPrimary [disabled]="!form.form.valid">
            <app-button-label icon="plus">
              Create
            </app-button-label>
          </button>
        </div>
      </div>
    </form>
  `
})
export class ProjectCreateComponent {
  @Output()
  projectCreated = new EventEmitter<CreateProjectPayload>();

  name: string;
  password: string;

  onSubmit() {
    this.projectCreated.emit({
      name: this.name,
      password: this.password
    });
  }
}
