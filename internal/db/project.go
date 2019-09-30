package db

type Project struct {
	Model

	Name         string `json:"name"`
	PasswordHash []byte `json:"-"`
}

func (db *DB) CreateProject(project Project) (Project, error) {
	err := db.Create(&project).Error
	return project, err
}

func (db *DB) LoadAllProjects() []Project {
	var projects []Project
	db.Find(&projects)
	return projects
}

func (db *DB) LoadProjectByName(name string) (Project, error) {
	var project Project
	err := db.Where(&Project{Name: name}).Find(&project).Error
	return project, err
}

func (db *DB) LoadProjectByID(id uint) (Project, error) {
	var project Project
	err := db.First(&project, id).Error
	return project, err
}
