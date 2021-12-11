import { Component, EventEmitter, Output } from '@angular/core';
import { MatchFormSettings } from '../services/match-form.service';

@Component({
  selector: 'app-match-settings',
  template: `
    <p class="lead">
      Change Match Settings
      <app-icon
        icon="times"
        class="float-right"
        [button]="true"
        (click)="onSettingsClose()"
      ></app-icon>
    </p>
    <form (ngSubmit)="onSubmit(); form.reset()" #form="ngForm">
      <input
        type="text"
        placeholder="Team Count"
        class="form-control"
        name="teamCount"
        required
        [(ngModel)]="teamCount"
      />
      <input
        type="text"
        placeholder="Team Size"
        class="form-control"
        name="teamSize"
        required
        [(ngModel)]="teamSize"
      />
      <div class="d-flex flex-row-reverse">
        <button type="submit" appPrimary [disabled]="!form.form.valid">
          <app-button-label icon="save"> Change </app-button-label>
        </button>
      </div>
    </form>
  `,
  styles: [
    `
      .form-control {
        margin-bottom: 8px;
      }
    `,
  ],
})
export class MatchSettingsComponent {
  @Output()
  formSettingsChanged = new EventEmitter<MatchFormSettings>();

  @Output()
  settingsClose = new EventEmitter();

  teamCount;
  teamSize;

  onSubmit() {
    this.formSettingsChanged.emit({
      teamCount: this.teamCount,
      teamSize: this.teamSize,
    });
    this.settingsClose.emit();
  }

  onSettingsClose() {
    this.settingsClose.emit();
  }
}
