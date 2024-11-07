import { Component, EventEmitter, Output } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { Player } from '@generated/api';

import { ButtonLabelComponent } from '../../../shared/button/button-label.component';
import { ButtonPrimaryDirective } from '../../../shared/button/button-primary.directive';

@Component({
  selector: 'app-player-create',
  template: `
    <p class="lead">Create Player</p>
    <form (ngSubmit)="onSubmit(); form.reset()" #form="ngForm">
      <div class="row">
        <div class="col-9">
          <input
            type="text"
            placeholder="Player Name"
            class="form-control"
            name="name"
            required
            [(ngModel)]="name"
          />
        </div>
        <div class="col-3 d-flex justify-content-end">
          <button type="submit" appPrimary [disabled]="!form.form.valid">
            <app-button-label icon="plus"> Create </app-button-label>
          </button>
        </div>
      </div>
    </form>
  `,
  standalone: true,
  imports: [FormsModule, ButtonPrimaryDirective, ButtonLabelComponent],
})
export class PlayerCreateComponent {
  @Output()
  playerCreated = new EventEmitter<Partial<Player>>();

  name: string;

  onSubmit() {
    this.playerCreated.emit({ name: this.name });
  }
}
