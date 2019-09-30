package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"compelo/internal/db"
)

const (
	ProjectID  = "projectID"
	ProjectKey = "project"
)

func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	if err := unmarshal(r.Body, &body); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	p, err := h.svc.CreateProject(body.Name, body.Password)
	if err == nil {
		writeJSON(w, http.StatusCreated, p)
	} else {
		writeError(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects := h.svc.LoadAllProjects()
	writeJSON(w, http.StatusOK, projects)
}

func (h *Handler) ProjectCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, ProjectID))
		if err != nil {
			writeError(w, http.StatusBadRequest, err)
			return
		}
		project, err := h.svc.LoadProjectByID(uint(id))
		if err != nil {
			msg := fmt.Sprintf("could not set project with id %d in context", id)
			writeError(w, http.StatusNotFound, fmt.Errorf("%s: %v", msg, err))
			return
		}
		ctx := context.WithValue(r.Context(), ProjectKey, project)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func mustLoadProjectFromContext(r *http.Request) db.Project {
	project, ok := r.Context().Value(ProjectKey).(db.Project)
	if !ok {
		panic("project must be set in context")
	}
	return project
}
