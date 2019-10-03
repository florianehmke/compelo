package db

type Project struct {
	Model

	Name         string `json:"name"`
	PasswordHash []byte `json:"-"`
}

func (db *DB) CreateProject(project Project) (Project, error) {
	err := db.gorm.Create(&project).Error
	return project, err
}

func (db *DB) LoadAllProjects() []Project {
	var projects []Project
	db.gorm.Find(&projects)
	return projects
}

func (db *DB) LoadProjectByName(name string) (Project, error) {
	var project Project
	err := db.gorm.Where(&Project{Name: name}).Find(&project).Error
	return project, err
}

func (db *DB) LoadProjectByID(id uint) (Project, error) {
	var project Project
	err := db.gorm.First(&project, id).Error
	return project, err
}
