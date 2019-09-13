import { createAction, props } from '@ngrx/store';
import { ErrorPayload, Game, Payload, Player } from '@shared/models';

export const loadGames = createAction('[Games] Load');
export const loadGamesSuccess = createAction(
  '[Games] Load Success',
  props<Payload<Game[]>>()
);
export const loadGamesError = createAction(
  '[Games] Load Error',
  props<ErrorPayload>()
);

export const createGame = createAction(
  '[Games] Create',
  props<Payload<Game>>()
);
export const createGameSuccess = createAction(
  '[Games] Create Success',
  props<Payload<Game>>()
);
export const createGameError = createAction(
  '[Games] Create Error',
  props<ErrorPayload>()
);

export const selectGame = createAction(
  '[Games] Select',
  props<Payload<Game>>()
);

export const loadPlayers = createAction('[Players] Load');
export const loadPlayersSuccess = createAction(
  '[Players] Load Success',
  props<Payload<Player[]>>()
);
export const loadPlayersError = createAction(
  '[Players] Load Error',
  props<ErrorPayload>()
);

export const createPlayer = createAction(
  '[Players] Create',
  props<Payload<Player>>()
);
export const createPlayerSuccess = createAction(
  '[Players] Create Success',
  props<Payload<Player>>()
);
export const createPlayerError = createAction(
  '[Players] Create Error',
  props<ErrorPayload>()
);
