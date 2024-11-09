import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';

import { ButtonLabelComponent } from '../../../shared/button/button-label.component';
import { ButtonPrimaryDirective } from '../../../shared/button/button-primary.directive';

@Component({
  selector: 'app-project-select-modal',
  template: `
    <form (ngSubmit)="onSubmit()" #form="ngForm">
      <div class="modal-header">
        <h4 class="modal-title">Enter Password</h4>
        <button
          type="button"
          class="btn-close"
          aria-label="Close"
          (click)="activeModal.dismiss('Cross click')"
        ></button>
      </div>
      <div class="modal-body">
        <div class="form-group">
          <label for="exampleInputPassword1">Password</label>
          <input
            name="password"
            type="password"
            class="form-control"
            ngbAutoFocus
            [(ngModel)]="password"
          />
        </div>
      </div>
      <div class="modal-footer">
        <button type="submit" appPrimary>
          <app-button-label icon="sign-in-alt"> Login</app-button-label>
        </button>
      </div>
    </form>
  `,
  standalone: true,
  imports: [FormsModule, ButtonPrimaryDirective, ButtonLabelComponent],
})
export class ProjectSelectModalComponent {
  password: string;

  constructor(public activeModal: NgbActiveModal) {}

  onSubmit() {
    if (this.password) {
      this.activeModal.close(this.password);
    } else {
      this.activeModal.dismiss('no password entered');
    }
  }
}
