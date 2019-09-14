package match

import (
	"compelo"
	"compelo/db"
	"compelo/game"
	"compelo/player"
	"time"
)

type Service struct {
	db *db.DB

	playerService *player.Service
	gameService   *game.Service
}

func NewService(
	db *db.DB,
	playerService *player.Service,
	gameService *game.Service,
) *Service {
	return &Service{
		db:            db,
		playerService: playerService,
		gameService:   gameService,
	}
}

type CreateMatchParameter struct {
	Date   time.Time
	GameID uint

	Teams         uint
	WinningTeam   uint
	TeamPlayerMap map[uint]uint
	TeamScoreMap  map[uint]int
}

// TODO wrap in txn
func (s *Service) CreateMatch(param CreateMatchParameter) (*compelo.Match, error) {
	// Verify that the game exists.
	g, err := s.gameService.LoadGameByID(param.GameID)
	if err != nil {
		return nil, err // FIXME
	}

	// Create the match.
	m, err := s.createMatch(param.Date, g.ID)
	if err != nil {
		return nil, err // FIXME
	}

	// Create the teams.
	teamMap := map[uint]*compelo.MatchTeam{}
	for i := 1; i <= int(param.Teams); i++ {
		t, err := s.createTeam(m.ID, param.TeamScoreMap[uint(i)]) // FIXME maybe validate scores
		if err != nil {
			return nil, err // FIXME
		}
		teamMap[uint(i)] = t
	}

	// Create the players.
	playerIdMatchPlayerMap := map[uint]*compelo.MatchPlayer{}
	for teamNumber, playerID := range param.TeamPlayerMap {
		p, err := s.playerService.LoadPlayerByID(playerID)
		if err != nil {
			return nil, err // FIXME
		}
		mp, err := s.createPlayer(m.ID, teamMap[teamNumber].ID, p.ID)
		if err != nil {
			return nil, err // FIXME
		}
		playerIdMatchPlayerMap[playerID] = mp
	}

	// Set the "real" match id of the winning team.
	m.WinnerMatchTeamID = teamMap[param.WinningTeam].ID
	err = s.updateMatch(m)
	return m, err
}

func (s *Service) createMatch(date time.Time, gameID uint) (*compelo.Match, error) {
	m := &compelo.Match{Date: date, GameID: gameID}
	err := s.db.Create(m).Error
	return m, err
}

func (s *Service) updateMatch(match *compelo.Match) error {
	return s.db.Save(match).Error
}

func (s *Service) createTeam(matchID uint, score int) (*compelo.MatchTeam, error) {
	t := &compelo.MatchTeam{MatchID: matchID, Score: score}
	err := s.db.Create(t).Error
	return t, err
}

func (s *Service) createPlayer(matchID uint, matchTeamID uint, playerID uint) (*compelo.MatchPlayer, error) {
	p := &compelo.MatchPlayer{MatchID: matchID, MatchTeamID: matchTeamID, PlayerID: playerID}
	err := s.db.Create(p).Error
	return p, err
}

func (s *Service) LoadMatches() []compelo.Match {
	var matches []compelo.Match
	s.db.Find(&matches)
	return matches
}

type CompleteMatch struct {
	Match        compelo.Match         `json:"match"`
	MatchTeams   []compelo.MatchTeam   `json:"teams"`
	MatchPlayers []compelo.MatchPlayer `json:"players"`
}

func (s *Service) LoadByID(id uint) (CompleteMatch, error) {
	var match compelo.Match
	err := s.db.First(&match, id).Error
	if err != nil {
		return CompleteMatch{}, err // FIXME
	}

	var players []compelo.MatchPlayer
	s.db.Where(&compelo.MatchPlayer{MatchID: match.ID}).Find(&players)

	var teams []compelo.MatchTeam
	s.db.Where(&compelo.MatchTeam{MatchID: match.ID}).Find(&teams)

	return CompleteMatch{
		Match:        match,
		MatchTeams:   teams,
		MatchPlayers: players,
	}, nil
}
