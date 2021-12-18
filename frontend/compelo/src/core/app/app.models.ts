export interface LoadingActions {
  [key: string]: boolean;
}

export const DEFAULT_ACTION_TYPES = ['@ngrx'];
export const SUCCESS_ACTION_TYPE = 'Success';
export const ERROR_ACTION_TYPE = 'Error';
