CREATE TABLE projects
(
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at    DATETIME,
    updated_at    DATETIME,
    deleted_at    DATETIME,
    name          VARCHAR(255) NOT NULL UNIQUE,
    password_hash BLOB         NOT NULL
);

CREATE INDEX idx_projects_deleted_at
    ON projects (deleted_at);



CREATE TABLE games
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    name       VARCHAR(255) NOT NULL UNIQUE,
    project_id INTEGER      NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE
);

CREATE INDEX idx_games_deleted_at
    ON games (deleted_at);



CREATE TABLE players
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    name       VARCHAR(255) NOT NULL UNIQUE,
    project_id INTEGER      NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE
);

CREATE INDEX idx_players_deleted_at
    ON players (deleted_at);



CREATE TABLE ratings
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    rating     INTEGER NOT NULL,
    game_id    INTEGER NOT NULL,
    player_id  INTEGER NOT NULL,
    FOREIGN KEY (game_id) REFERENCES games (id) ON DELETE CASCADE,
    FOREIGN KEY (player_id) REFERENCES players (id)
);

CREATE INDEX idx_ratings_deleted_at
    ON ratings (deleted_at);



CREATE TABLE matches
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    date       DATETIME NOT NULL,
    game_id    INTEGER,
    FOREIGN KEY (game_id) REFERENCES games (id) ON DELETE CASCADE
);

CREATE INDEX idx_matches_deleted_at
    ON matches (deleted_at);



CREATE TABLE appearances
(
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at   DATETIME,
    updated_at   DATETIME,
    deleted_at   DATETIME,
    match_id     INTEGER NOT NULL,
    team_id      INTEGER NOT NULL,
    player_id    INTEGER NOT NULL,
    rating_delta INTEGER NOT NULL,
    FOREIGN KEY (match_id) REFERENCES matches (id) ON DELETE CASCADE,
    FOREIGN KEY (team_id) REFERENCES teams (id) ON DELETE CASCADE,
    FOREIGN KEY (player_id) REFERENCES players (id)
);

CREATE INDEX idx_appearances_deleted_at
    ON appearances (deleted_at);



CREATE TABLE teams
(
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at   DATETIME,
    updated_at   DATETIME,
    deleted_at   DATETIME,
    match_id     INTEGER      NOT NULL,
    score        INTEGER      NOT NULL,
    result       VARCHAR(255) NOT NULL,
    rating_delta INTEGER      NOT NULL,
    FOREIGN KEY (match_id) REFERENCES matches (id) ON DELETE CASCADE
);

CREATE INDEX idx_teams_deleted_at
    ON teams (deleted_at);



CREATE VIEW match_results AS
SELECT a.player_id, m.game_id, m.date, t.rating_delta, t.result
FROM matches m
         JOIN appearances a ON m.id = a.match_id
         JOIN teams t ON a.team_id = t.id
ORDER BY m.date;
