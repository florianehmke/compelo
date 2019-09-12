import { Component } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';

@Component({
  selector: 'app-project-select-modal',
  template: `
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
        <input type="password" class="form-control" [(ngModel)]="password" />
      </div>
    </div>
    <div class="modal-footer">
      <button
        type="button"
        class="btn btn-primary"
        aria-label="Confirm"
        (click)="onConfirm()"
      >
        Confirm
      </button>
    </div>
  `
})
export class ProjectSelectModalComponent {
  password: string;

  constructor(public activeModal: NgbActiveModal) {}

  onConfirm() {
    if (this.password) {
      this.activeModal.close(this.password);
    } else {
      this.activeModal.dismiss('no password entered');
    }
  }
}
