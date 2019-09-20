package player

import (
	"compelo/db"
	"compelo/rating"
)

type Service struct {
	db *db.DB
}

func NewService(db *db.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreatePlayer(projectID uint, name string) (Player, error) {
	p := Player{
		Name:      name,
		ProjectID: projectID,
	}
	err := s.db.Create(&p).Error
	return p, err
}

func (s *Service) LoadPlayersByProjectID(pid uint) ([]Player, error) {
	var players []Player
	err := s.db.Where(&Player{ProjectID: pid}).Order("id").Find(&players).Error
	return players, err
}

func (s *Service) LoadPlayerByID(id uint) (Player, error) {
	var player Player
	err := s.db.First(&player, id).Error
	return player, err
}

func (s *Service) LoadRating(playerID, gameID uint) (Rating, error) {
	var r Rating
	ref := s.db.Where(Rating{GameID: gameID, PlayerID: playerID}).First(&r)
	if ref.RecordNotFound() {
		r = Rating{
			Rating:   rating.InitialRating,
			GameID:   gameID,
			PlayerID: playerID,
		}
		s.db.Create(&r)
		return r, nil
	}
	return r, ref.Error
}

func (s *Service) UpdateRating(playerID, gameID uint, delta int) (Rating, error) {
	r, err := s.LoadRating(playerID, gameID)
	if err != nil {
		return Rating{}, nil
	}

	r.Rating = r.Rating + delta
	err = s.db.Save(&r).Error
	return r, err
}
