import { NgIf, NgFor } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';
import {
  UntypedFormGroup,
  FormsModule,
  ReactiveFormsModule,
} from '@angular/forms';

import { CreateMatchRequest, Player } from '@generated/api';

import { ButtonLabelComponent } from '../../../shared/button/button-label.component';
import { ButtonPrimaryDirective } from '../../../shared/button/button-primary.directive';
import { IconComponent } from '../../../shared/icon/icon.component';

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
          [formGroupName]="i.toString()"
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
              [formControlName]="j.toString()"
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
        <button type="submit" appPrimary [disabled]="!formGroup.valid">
          <app-button-label icon="plus"> Submit </app-button-label>
        </button>
      </div>
    </form>
  `,
  styles: [
    `
      .form-control,
      .custom-select {
        margin-bottom: 8px;
      }
    `,
  ],
  standalone: true,
  imports: [
    IconComponent,
    NgIf,
    FormsModule,
    ReactiveFormsModule,
    NgFor,
    ButtonPrimaryDirective,
    ButtonLabelComponent,
  ],
})
export class MatchCreateComponent {
  @Input()
  players: Player[];

  @Input()
  formGroup: UntypedFormGroup;

  @Output()
  matchCreated = new EventEmitter<CreateMatchRequest>();

  @Output()
  settingsClick = new EventEmitter();

  onSubmit() {
    if (this.formGroup.valid) {
      const value: CreateMatchRequest = this.formGroup.value;
      this.matchCreated.emit(value);
      this.formGroup.reset();
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
