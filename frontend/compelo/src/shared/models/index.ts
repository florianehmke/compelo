export * from './project.models';
export * from './game.models';
export * from './player.models';

export interface Payload<T> {
  payload: T;
}

export type ErrorPayload = Payload<any>;
