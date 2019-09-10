export * from './project.models';

export interface Payload<T> {
  payload: T;
}

export type ErrorPayload = Payload<any>;
