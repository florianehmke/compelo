package handler

import (
	"net/http"

	"compelo/api/json"
	"compelo/command"
)

const (
	ProjectGUID string = "projectGUID"
)

type CreateProjectRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var request CreateProjectRequest
	if err := json.Unmarshal(r.Body, &request); err != nil {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	p, err := h.c.CreateNewProject(command.CreateNewProjectCommand{
		Name:     request.Name,
		Password: request.Password,
	})
	if err == nil {
		json.WriteResponse(w, http.StatusCreated, p)
	} else {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects := h.q.GetProjects()
	json.WriteResponse(w, http.StatusOK, projects)
}
