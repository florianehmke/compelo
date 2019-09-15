import { Component, EventEmitter, Output } from '@angular/core';
import { Game } from '@shared/models';
import { CreateProjectPayload } from '@core/project-list';

@Component({
  selector: 'app-game-create',
  template: `
    <form (ngSubmit)="onSubmit(form.value); form.reset()" #form="ngForm">
      <div class="row">
        <div class="col-8">
          <input
            type="text"
            placeholder="Game Name"
            class="form-control"
            name="name"
            required
            [(ngModel)]="name"
          />
        </div>
        <div class="col-4">
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
export class GameCreateComponent {
  @Output()
  gameCreated = new EventEmitter<Game>();

  name: string;

  onSubmit(value: CreateProjectPayload) {
    this.gameCreated.emit(value);
  }
}
