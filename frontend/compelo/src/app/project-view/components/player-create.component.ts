import { Component, EventEmitter, Output } from '@angular/core';
import { Player } from '@shared/models';
import { CreateProjectPayload } from '@core/project-list';

@Component({
  selector: 'app-player-create',
  template: `
    <p class="lead">Create Player</p>
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
          <app-button
            class="w-100"
            type="submit"
            icon="plus"
            [disabled]="!form.form.valid"
            label="Create"
          >
          </app-button>
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
