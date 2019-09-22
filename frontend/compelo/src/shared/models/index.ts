export * from './compelo.models';
export * from './auth.models';

export interface Payload<T> {
  payload: T;
}

export type ErrorPayload = Payload<any>;
