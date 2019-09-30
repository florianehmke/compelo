package db

type Game struct {
	Model

	Name      string `json:"name"`
	ProjectID uint   `json:"projectId"`
}

func (db *DB) CreateGame(game Game) (Game, error) {
	err := db.Create(&game).Error
	return game, err
}

func (db *DB) LoadGamesByProjectID(projectID uint) []Game {
	var games []Game
	db.Where(&Game{ProjectID: projectID}).Find(&games)
	return games
}

func (db *DB) LoadGameByID(id uint) (Game, error) {
	var game Game
	err := db.First(&game, id).Error
	return game, err
}