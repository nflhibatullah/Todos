package project

import (
	"todos/entities"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (prrep *ProjectRepository) GetAll() ([]entities.Project, error) {
	project := []entities.Project{}
	prrep.db.Find(&project)
	return project, nil
}

func (prrep *ProjectRepository) Get(projectId int) (entities.Project, error) {
	project := entities.Project{}
	prrep.db.Find(&project, projectId)
	return project, nil
}

func (prrep *ProjectRepository) Create(project entities.Project) (entities.Project, error) {
	prrep.db.Save(&project)
	return project, nil
}

func (prrep *ProjectRepository) Delete(projectId int) (entities.Project, error) {
	project := entities.Project{}
	prrep.db.Find(&project, "id=?", projectId)
	prrep.db.Delete(&project)
	return project, nil
}

func (prrep *ProjectRepository) Update(newProject entities.Project, projectId int) (entities.Project, error) {
	project := entities.Project{}
	prrep.db.Find(&project, "id=?", projectId)
	prrep.db.Model(&project).Updates(newProject)
	return newProject, nil
}
