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



CREATE TABLE match_players
(
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at    DATETIME,
    updated_at    DATETIME,
    deleted_at    DATETIME,
    match_id      INTEGER NOT NULL,
    match_team_id INTEGER NOT NULL,
    player_id     INTEGER NOT NULL,
    FOREIGN KEY (match_id) REFERENCES matches (id) ON DELETE CASCADE,
    FOREIGN KEY (match_team_id) REFERENCES match_teams (id) ON DELETE CASCADE,
    FOREIGN KEY (player_id) REFERENCES players (id)
);

CREATE INDEX idx_match_players_deleted_at
    ON match_players (deleted_at);



CREATE TABLE match_teams
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    match_id   INTEGER NOT NULL,
    score      INTEGER,
    winner     INTEGER,
    FOREIGN KEY (match_id) REFERENCES matches (id) ON DELETE CASCADE
);

CREATE INDEX idx_match_teams_deleted_at
    ON match_teams (deleted_at);


