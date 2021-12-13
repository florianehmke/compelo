import { Spectator, createComponentFactory } from '@ngneat/spectator/jest';

import { StatsBarComponent } from './stats-bar.component';

describe('StatsBarComponent', () => {
  let spectator: Spectator<StatsBarComponent>;
  const createComponent = createComponentFactory(StatsBarComponent);

  beforeEach(() => (spectator = createComponent()));

  it('should not render segments if data input is null', () => {
    expect(spectator.query('.segment')).toBeNull();
  });

  it('should render segments according to the input data', () => {
    spectator.setInput('data', {
      wins: 3,
      draws: 2,
      lost: 1,
    });

    expect(spectator.query('.segment.wins')).toHaveText('3');
    expect(spectator.query('.segment.draws')).toHaveText('2');
    expect(spectator.query('.segment.lost')).toHaveText('1');
  });
});
