import { Component, EventEmitter, Input, Output } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { Player } from '@shared/models';
import { CreateMatchPayload } from '@core/project';

@Component({
  selector: 'app-match-create',
  template: `
    <p class="lead">
      Create Match
      <app-icon
        icon="wrench"
        class="float-right"
        [button]="true"
        (click)="onSettingsClick()"
      ></app-icon>
    </p>
    <form [formGroup]="formGroup" (ngSubmit)="onSubmit()" *ngIf="players">
      <div class="row" formArrayName="teams">
        <div
          class="col"
          [formGroupName]="i"
          *ngFor="let team of getTeams(formGroup); let i = index"
        >
          <input
            type="number"
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
      <div class="d-flex flex-row-reverse">
        <app-button
          type="submit"
          icon="plus"
          [disabled]="!formGroup.valid"
          label="Submit"
        >
        </app-button>
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

  @Output()
  matchCreated = new EventEmitter<CreateMatchPayload>();

  @Output()
  settingsClick = new EventEmitter();

  onSubmit() {
    if (this.formGroup.valid) {
      const value: CreateMatchPayload = this.formGroup.value;
      this.matchCreated.emit(value);
    }
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

  onSettingsClick() {
    this.settingsClick.emit();
  }
}
