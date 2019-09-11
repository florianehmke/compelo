import { Component, EventEmitter, Output } from '@angular/core';
import { CreateProjectPayload } from '../../../shared/models';

@Component({
  selector: 'app-project-create',
  template: `
    <form (ngSubmit)="onSubmit(form.value); form.reset()" #form="ngForm">
      <div class="row">
        <div class="col-5">
          <input
            type="text"
            placeholder="Project Name"
            class="form-control"
            name="name"
            required
            [(ngModel)]="name"
          />
        </div>
        <div class="col-5">
          <input
            type="password"
            placeholder="Password"
            class="form-control"
            name="password"
            required
            [(ngModel)]="password"
          />
        </div>
        <div class="col-2">
          <button
            type="submit"
            class="w-100 btn btn-success"
            [class.disabled]="!form.form.valid"
            [disabled]="!form.form.valid"
          >
            Create
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

  onSubmit(value: CreateProjectPayload) {
    console.log(value);
    this.projectCreated.emit(value);
  }
}
