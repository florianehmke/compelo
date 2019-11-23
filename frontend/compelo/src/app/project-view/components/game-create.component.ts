import { Component, EventEmitter, Output } from '@angular/core';
import { Game } from '@api';

@Component({
  selector: 'app-game-create',
  template: `
    <p class="lead">
      Create / Select Game
    </p>
    <form (ngSubmit)="onSubmit(); form.reset()" #form="ngForm">
      <div class="row">
        <div class="col-9">
          <input
            type="text"
            placeholder="Game Name"
            class="form-control"
            name="name"
            required
            [(ngModel)]="name"
          />
        </div>
        <div class="col-3 d-flex justify-content-end">
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
export class GameCreateComponent {
  @Output()
  gameCreated = new EventEmitter<Partial<Game>>();

  name: string;

  onSubmit() {
    this.gameCreated.emit({ name: this.name });
  }
}
