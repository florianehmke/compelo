import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs';
import { flatMap, map, take } from 'rxjs/operators';

import { environment } from '@env/environment';
import {
  Competition,
  CreateCompetitionRequest,
  CreateMatchRequest,
  Game,
  GameStats,
  Match,
  Player,
  PlayerStats,
  Response,
} from '@generated/api';

import { getSelectedProjectGuid } from '../router';
import { State } from '../router/router-state.reducer';

@Injectable()
export class ProjectService {
  baseUrl = environment.baseUrl + '/projects';

  constructor(private http: HttpClient, private routerStore: Store<State>) {}

  getGames(): Observable<Game[]> {
    return this.projectUrl().pipe(
      flatMap((url) => this.http.get<Game[]>(`${url}/games`))
    );
  }

  createGame(game: Game): Observable<Response> {
    return this.projectUrl().pipe(
      flatMap((url) => this.http.post<Response>(`${url}/games`, game))
    );
  }

  getGameStats(gameGuid: string): Observable<GameStats> {
    return this.projectUrl().pipe(
      flatMap((url) =>
        this.http.get<GameStats>(`${url}/games/${gameGuid}/game-stats`)
      )
    );
  }

  getPlayers(): Observable<Player[]> {
    return this.projectUrl().pipe(
      flatMap((url) => this.http.get<Player[]>(`${url}/players`))
    );
  }

  createPlayer(game: Player): Observable<Response> {
    return this.projectUrl().pipe(
      flatMap((url) => this.http.post<Response>(`${url}/players`, game))
    );
  }

  createMatch(
    payload: CreateMatchRequest,
    gameGuid: string
  ): Observable<Response> {
    return this.projectUrl().pipe(
      flatMap((url) =>
        this.http.post<Response>(`${url}/games/${gameGuid}/matches`, payload)
      )
    );
  }

  deleteMatch(matchGuid: string, gameGuid: string): Observable<Response> {
    return this.projectUrl().pipe(
      flatMap((url) =>
        this.http.delete<Response>(
          `${url}/games/${gameGuid}/matches/${matchGuid}`
        )
      )
    );
  }

  getMatches(gameGuid: string): Observable<Match[]> {
    return this.projectUrl().pipe(
      flatMap((url) =>
        this.http.get<Match[]>(`${url}/games/${gameGuid}/matches`)
      )
    );
  }

  getPlayerStats(gameGuid: string): Observable<PlayerStats[]> {
    return this.projectUrl().pipe(
      flatMap((url) =>
        this.http.get<PlayerStats[]>(`${url}/games/${gameGuid}/player-stats`)
      )
    );
  }

  createCompetition(
    payload: CreateCompetitionRequest,
    gameGuid: string
  ): Observable<Response> {
    return this.projectUrl().pipe(
      flatMap((url) =>
        this.http.post<Response>(
          `${url}/games/${gameGuid}/competitions`,
          payload
        )
      )
    );
  }

  getCompetitions(gameGuid: string): Observable<Competition[]> {
    return this.projectUrl().pipe(
      flatMap((url) =>
        this.http.get<Competition[]>(`${url}/games/${gameGuid}/competitions`)
      )
    );
  }

  projectUrl(): Observable<string> {
    return this.routerStore.select(getSelectedProjectGuid).pipe(
      take(1),
      map((id) => `${this.baseUrl}/${id}`)
    );
  }
}
