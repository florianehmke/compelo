package query

import (
	"errors"
	"fmt"
)

var ErrProjectNotFound = errors.New("project not found")

func (svc *Service) GetProjects() []*Project {
	svc.RLock()
	defer svc.RUnlock()

	list := make([]*Project, 0, len(svc.data.projects))
	for _, value := range svc.data.projects {
		list = append(list, value)
	}

	sortProjectsByCreatedDate(list)
	return list
}

func (svc *Service) GetProjectBy(projectGUID string) (*Project, error) {
	svc.RLock()
	defer svc.RUnlock()

	return svc.getProjectBy(projectGUID)
}

func (svc *Service) getProjectBy(projectGUID string) (*Project, error) {
	if project, ok := svc.data.projects[projectGUID]; ok {
		return project, nil
	}
	return nil, fmt.Errorf("get project by guid (%s) failed: %w", projectGUID, ErrProjectNotFound)
}
