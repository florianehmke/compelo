package db

type Player struct {
	Model

	Name      string `json:"name"`
	ProjectID uint   `json:"projectId"`
}

func (db *gormDB) CreatePlayer(player Player) (Player, error) {
	err := db.gorm.Create(&player).Error
	return player, err
}

func (db *gormDB) LoadPlayerByID(id uint) (Player, error) {
	var player Player
	err := db.gorm.First(&player, id).Error
	return player, err
}

func (db *gormDB) LoadPlayersByProjectID(projectID uint) []Player {
	var players []Player
	db.gorm.Where(&Player{ProjectID: projectID}).Find(&players)
	return players
}

func (db *gormDB) LoadPlayersByMatchIDAndTeamID(matchID, teamID uint) ([]Player, error) {
	var players []Player
	err := db.gorm.
		Joins("left join appearances on appearances.player_id = players.id").
		Where("appearances.match_id = ? and appearances.team_id = ? ", matchID, teamID).
		Find(&players).Error

	return players, err
}
