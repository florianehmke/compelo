package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"compelo/api/json"
	"compelo/command"
	"compelo/query"
)

const (
	ProjectGUID = "projectGUID"
	ProjectKey  = "project"
)

func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var body command.CreateNewProjectCommand
	if err := json.Unmarshal(r.Body, &body); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	p, err := h.c.CreateNewProject(body)
	if err == nil {
		json.Write(w, http.StatusCreated, p)
	} else {
		json.WriteError(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects := h.q.GetProjects()
	json.Write(w, http.StatusOK, projects)
}

func (h *Handler) ProjectCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		guid := chi.URLParam(r, ProjectGUID)
		if guid == "" {
			json.WriteError(w, http.StatusBadRequest, errors.New("no guid provided"))
		}
		project, err := h.q.GetProjectBy(chi.URLParam(r, ProjectGUID))
		if err != nil {
			msg := fmt.Sprintf("could not set project with guid %s in context", guid)
			json.WriteError(w, http.StatusNotFound, fmt.Errorf("%s: %v", msg, err))
			return
		}
		ctx := context.WithValue(r.Context(), ProjectKey, project)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func MustLoadProjectFromContext(r *http.Request) query.Project {
	project, ok := r.Context().Value(ProjectKey).(query.Project)
	if !ok {
		panic("project must be set in context")
	}
	return project
}
