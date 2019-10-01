import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '@env/environment';
import { Game, Match, Player, PlayerStats } from '@shared/models';
import { Store } from '@ngrx/store';
import { flatMap, map, take } from 'rxjs/operators';

import { CreateMatchPayload } from './project.models';
import { State } from '../router/router-state.reducer';
import { getSelectedProjectId } from '../router';

@Injectable()
export class ProjectService {
  baseUrl = environment.baseUrl + '/projects';

  constructor(private http: HttpClient, private routerStore: Store<State>) {}

  getGames(): Observable<Game[]> {
    return this.projectUrl().pipe(
      flatMap(url => this.http.get<Game[]>(`${url}/games`))
    );
  }

  createGame(game: Game): Observable<Game> {
    return this.projectUrl().pipe(
      flatMap(url => this.http.post<Game>(`${url}/games`, game))
    );
  }

  getPlayers(): Observable<Player[]> {
    return this.projectUrl().pipe(
      flatMap(url => this.http.get<Player[]>(`${url}/players`))
    );
  }

  createPlayer(game: Player): Observable<Player> {
    return this.projectUrl().pipe(
      flatMap(url => this.http.post<Player>(`${url}/players`, game))
    );
  }

  createMatch(payload: CreateMatchPayload, gameID: number): Observable<Match> {
    return this.projectUrl().pipe(
      flatMap(url =>
        this.http.post<Match>(`${url}/games/${gameID}/matches`, payload)
      )
    );
  }

  getMatches(gameID: number): Observable<Match[]> {
    return this.projectUrl().pipe(
      flatMap(url => this.http.get<Match[]>(`${url}/games/${gameID}/matches`))
    );
  }

  getPlayersWithStats(gameID: number): Observable<PlayerStats[]> {
    return this.projectUrl().pipe(
      flatMap(url =>
        this.http.get<PlayerStats[]>(`${url}/games/${gameID}/player-stats`)
      )
    );
  }

  projectUrl(): Observable<string> {
    return this.routerStore.select(getSelectedProjectId).pipe(
      take(1),
      map(id => `${this.baseUrl}/${id}`)
    );
  }
}
