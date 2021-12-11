import { Component } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';

@Component({
  selector: 'app-project-select-modal',
  template: `
    <form (ngSubmit)="onSubmit()" #form="ngForm">
      <div class="modal-header">
        <h4 class="modal-title">Enter Password</h4>
        <button
          type="button"
          class="close"
          aria-label="Close"
          (click)="activeModal.dismiss('Cross click')"
        >
          <span aria-hidden="true">&times;</span>
        </button>
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
          <app-button-label icon="sign-in-alt"> Login </app-button-label>
        </button>
      </div>
    </form>
  `,
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
