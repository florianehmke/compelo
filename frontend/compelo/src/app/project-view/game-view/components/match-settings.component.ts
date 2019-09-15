import { Component, EventEmitter, Output } from '@angular/core';
import { MatchFormSettings } from '../services/match-form.service';

@Component({
  selector: 'app-match-settings',
  template: `
    <form (ngSubmit)="onSubmit(); form.reset()" #form="ngForm">
      <div class="row">
        <div class="col-4">
          <input
            type="text"
            placeholder="Team Count"
            class="form-control"
            name="teamCount"
            required
            [(ngModel)]="teamCount"
          />
        </div>
        <div class="col-4">
          <input
            type="text"
            placeholder="Team Size"
            class="form-control"
            name="teamSize"
            required
            [(ngModel)]="teamSize"
          />
        </div>
        <div class="col-4">
          <button
            type="submit"
            class="w-100 btn btn-primary"
            [class.disabled]="!form.form.valid"
            [disabled]="!form.form.valid"
          >
            Change
          </button>
        </div>
      </div>
    </form>
  `
})
export class MatchSettingsComponent {
  @Output()
  formSettingsChanged = new EventEmitter<MatchFormSettings>();

  teamCount;
  teamSize;

  onSubmit() {
    this.formSettingsChanged.emit({
      teamCount: this.teamCount,
      teamSize: this.teamSize
    });
  }
}
