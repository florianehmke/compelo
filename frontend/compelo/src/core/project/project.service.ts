import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs';
import { flatMap, map, take } from 'rxjs/operators';

import { environment } from '@env/environment';
import {
  CreateMatchRequest,
  Game,
  GameStats,
  Match,
  MatchData,
  Player,
  PlayerStats,
} from '@generated/api';

import { getSelectedProjectId } from '../router';
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

  createGame(game: Game): Observable<Game> {
    return this.projectUrl().pipe(
      flatMap((url) => this.http.post<Game>(`${url}/games`, game))
    );
  }

  getGameStats(gameID: number): Observable<GameStats> {
    return this.projectUrl().pipe(
      flatMap((url) =>
        this.http.get<GameStats>(`${url}/games/${gameID}/game-stats`)
      )
    );
  }

  getPlayers(): Observable<Player[]> {
    return this.projectUrl().pipe(
      flatMap((url) => this.http.get<Player[]>(`${url}/players`))
    );
  }

  createPlayer(game: Player): Observable<Player> {
    return this.projectUrl().pipe(
      flatMap((url) => this.http.post<Player>(`${url}/players`, game))
    );
  }

  createMatch(payload: CreateMatchRequest, gameID: number): Observable<Match> {
    return this.projectUrl().pipe(
      flatMap((url) =>
        this.http.post<Match>(`${url}/games/${gameID}/matches`, payload)
      )
    );
  }

  getMatches(gameID: number): Observable<MatchData[]> {
    return this.projectUrl().pipe(
      flatMap((url) =>
        this.http.get<MatchData[]>(`${url}/games/${gameID}/matches`)
      )
    );
  }

  getPlayerStats(gameID: number): Observable<PlayerStats[]> {
    return this.projectUrl().pipe(
      flatMap((url) =>
        this.http.get<PlayerStats[]>(`${url}/games/${gameID}/player-stats`)
      )
    );
  }

  projectUrl(): Observable<string> {
    return this.routerStore.select(getSelectedProjectId).pipe(
      take(1),
      map((id) => `${this.baseUrl}/${id}`)
    );
  }
}
