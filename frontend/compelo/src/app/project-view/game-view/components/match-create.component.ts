import { Component, Input } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { Player } from '@shared/models';

@Component({
  selector: 'app-match-create',
  template: `
    <form [formGroup]="formGroup" (ngSubmit)="onSubmit()" *ngIf="players">
      <div class="row" formArrayName="teams">
        <div
          class="col"
          [formGroupName]="i"
          *ngFor="let team of getTeams(formGroup); let i = index"
        >
          <input
            type="text"
            placeholder="Score"
            class="form-control"
            formControlName="score"
          />
          <div formArrayName="playerIds">
            <select
              class="custom-select"
              [formControlName]="j"
              [compareWith]="compareByID"
              *ngFor="let player of getPlayers(team); let j = index"
            >
              <option *ngFor="let p of players" [ngValue]="p?.id">
                {{ p?.name }}
              </option>
            </select>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col-4">
          <button type="submit" class="w-100 btn btn-success">
            Submit
          </button>
        </div>
      </div>
    </form>
  `,
  styles: [
    `
      .form-control,
      .custom-select {
        margin-bottom: 8px;
      }
    `
  ]
})
export class MatchCreateComponent {

  @Input()
  players: Player[];

  @Input()
  formGroup: FormGroup;

  onSubmit() {
    // Determine winner
    // Validate

    console.log(this.formGroup.value);
    console.log(JSON.stringify(this.formGroup.value))
  }

  getTeams(form) {
    return form.controls.teams.controls;
  }

  getPlayers(form) {
    return form.controls.playerIds.controls;
  }

  compareByID(p1, p2): boolean {
    return p1 && p2 && p1.id === p2.id;
  }
}
