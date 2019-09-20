import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '@env/environment';
import { Game, Match, Player } from '@shared/models';
import { CreateMatchPayload } from '@core/project/project.models';

@Injectable()
export class ProjectService {
  private baseUrl = environment.baseUrl + '/project';

  constructor(private http: HttpClient) {}

  getGames(): Observable<Game[]> {
    return this.http.get<Game[]>(`${this.baseUrl}/games`);
  }

  createGame(game: Game): Observable<Game> {
    return this.http.post<Game>(`${this.baseUrl}/games`, game);
  }

  getPlayers(): Observable<Player[]> {
    return this.http.get<Player[]>(`${this.baseUrl}/players`);
  }

  createPlayer(game: Player): Observable<Player> {
    return this.http.post<Player>(`${this.baseUrl}/players`, game);
  }

  createMatch(payload: CreateMatchPayload, gameID: number): Observable<Match> {
    return this.http.post<Match>(`${this.baseUrl}/games/${gameID}/matches`, payload);
  }

  getMatches(gameID: number): Observable<Match[]> {
    return this.http.get<Match[]>(`${this.baseUrl}/games/${gameID}/matches`);
  }

  getPlayersWithStats(gameID: number): Observable<Player[]> {
    return this.http.get<Player[]>(`${this.baseUrl}/games/${gameID}/players`);
  }
}
