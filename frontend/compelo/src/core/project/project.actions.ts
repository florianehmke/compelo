import { createAction, props } from '@ngrx/store';

import {
  CreateMatchRequest,
  Game,
  GameStats,
  Match,
  Response,
  Player,
  PlayerStats,
} from '@generated/api';
import { ErrorPayload, Payload } from '@shared/models';

import {
  FilterMatchesPayload,
  LoadGameStatsPayload,
  LoadMatchesPayload,
  LoadPlayerStatsPayload,
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
  props<Payload<Response>>()
);
export const createGameError = createAction(
  '[Games] Create Error',
  props<ErrorPayload>()
);

export const loadGameStats = createAction(
  '[Games] Load Game Stats',
  props<Payload<LoadGameStatsPayload>>()
);
export const loadGameStatsSuccess = createAction(
  '[Games] Load Game Stats Success',
  props<Payload<GameStats>>()
);
export const loadGameStatsError = createAction(
  '[Games] Load Game Stats Error',
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
  props<Payload<Response>>()
);
export const createPlayerError = createAction(
  '[Players] Create Error',
  props<ErrorPayload>()
);

export const createMatch = createAction(
  '[Matches] Create Match',
  props<Payload<CreateMatchRequest>>()
);
export const createMatchSuccess = createAction(
  '[Matches] Create Match Success',
  props<Payload<Response>>()
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

export const loadPlayerStats = createAction(
  '[Players] Load Player Stats',
  props<Payload<LoadPlayerStatsPayload>>()
);
export const loadPlayerStatsSuccess = createAction(
  '[Players] Load Player Stats Success',
  props<Payload<PlayerStats[]>>()
);
export const loadPlayerStatsError = createAction(
  '[Players] Load Player Stats Error',
  props<ErrorPayload>()
);

export const filterMatches = createAction(
  '[Matches] Filter',
  props<Payload<FilterMatchesPayload>>()
);
