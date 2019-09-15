export * from './compelo.models';

export interface Payload<T> {
  payload: T;
}

export type ErrorPayload = Payload<any>;
