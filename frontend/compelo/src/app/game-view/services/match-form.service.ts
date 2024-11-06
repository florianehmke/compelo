import { Injectable } from '@angular/core';
import {
  UntypedFormBuilder,
  UntypedFormControl,
  UntypedFormGroup,
  Validators,
} from '@angular/forms';

export interface MatchFormSettings {
  teamCount: number;
  teamSize: number;
}

@Injectable()
export class MatchFormService {
  constructor(private fb: UntypedFormBuilder) {}

  buildForm(settings: MatchFormSettings): UntypedFormGroup {
    const teamArray = this.fb.array([]);
    for (let i = 0; i < settings.teamCount; i++) {
      teamArray.push(this.createTeamForm(settings.teamSize));
    }

    return this.fb.group({
      teams: teamArray,
    });
  }

  private createTeamForm(teamSize: number): UntypedFormGroup {
    const playerArray = this.fb.array([]);
    for (let i = 0; i < teamSize; i++) {
      playerArray.push(new UntypedFormControl(null, Validators.required));
    }

    return this.fb.group({
      playerIds: playerArray,
      score: [null, Validators.required],
    });
  }
}
