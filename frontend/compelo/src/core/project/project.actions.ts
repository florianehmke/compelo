import { createAction, props } from '@ngrx/store';
import { ErrorPayload, Game, Match, Payload, Player } from '@shared/models';
import {
  CreateMatchPayload,
  LoadMatchesPayload,
  LoadPlayersWithStatusPayload
} from './project.models';

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

export const createMatch = createAction(
  '[Matches] Create Match',
  props<Payload<CreateMatchPayload>>()
);
export const createMatchSuccess = createAction(
  '[Matches] Create Match Success',
  props<Payload<Match>>()
);
export const createMatchError = createAction(
  '[Matches] Create Match Error',
  props<ErrorPayload>()
);

export const loadMatches = createAction(
  '[Matches] Load',
  props<Payload<LoadMatchesPayload>>()
);
export const loadMatchesSuccess = createAction(
  '[Matches] Load Success',
  props<Payload<Match[]>>()
);
export const loadMatchesError = createAction(
  '[Matches] Load Error',
  props<ErrorPayload>()
);

export const loadPlayersWithStats = createAction(
  '[Players] Load With Stats',
  props<Payload<LoadPlayersWithStatusPayload>>()
);
export const loadPlayersWithStatsSuccess = createAction(
  '[Players] Load With Stats Success',
  props<Payload<Player[]>>()
);
export const loadPlayersWithStatsError = createAction(
  '[Players] Load With Stats Error',
  props<ErrorPayload>()
);
