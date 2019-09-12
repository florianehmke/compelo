import { Component, EventEmitter, Output } from '@angular/core';
import { CreateProjectPayload, Player } from '../../../shared/models';

@Component({
  selector: 'app-player-create',
  template: `
    <form (ngSubmit)="onSubmit(form.value); form.reset()" #form="ngForm">
      <div class="row">
        <div class="col-8">
          <input
            type="text"
            placeholder="Player Name"
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
export class PlayerCreateComponent {
  @Output()
  playerCreated = new EventEmitter<Player>();

  name: string;

  onSubmit(value: CreateProjectPayload) {
    this.playerCreated.emit(value);
  }
}
