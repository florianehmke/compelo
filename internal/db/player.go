package db

type Player struct {
	Model

	Name      string `json:"name"`
	ProjectID uint   `json:"projectId"`
}

func (db *DB) CreatePlayer(player Player) (Player, error) {
	err := db.gorm.Create(&player).Error
	return player, err
}

func (db *DB) LoadPlayersByProjectID(projectID uint) []Player {
	var players []Player
	db.gorm.Where(&Player{ProjectID: projectID}).Find(&players)
	return players
}

func (db *DB) LoadPlayerByID(id uint) (Player, error) {
	var player Player
	err := db.gorm.First(&player, id).Error
	return player, err
}
