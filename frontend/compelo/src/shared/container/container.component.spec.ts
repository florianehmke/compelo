import { createComponentFactory, Spectator } from '@ngneat/spectator/jest';

import { ContainerComponent } from './container.component';

describe('ContainerComponent', () => {
  let spectator: Spectator<ContainerComponent>;
  const createComponent = createComponentFactory(ContainerComponent);

  beforeEach(() => (spectator = createComponent()));

  it('should create', () => {
    expect(spectator.component).toBeTruthy();
  });
});
