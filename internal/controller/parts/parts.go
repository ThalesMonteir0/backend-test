package parts

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ThalesMonteir0/backend-test/internal/DTO"
)

type createProcessor interface {
	Execute(ctx context.Context, in DTO.Part) (DTO.Part, error)
}

type getProcessor interface {
	Execute(ctx context.Context) ([]DTO.Part, error)
}

type deleteProcessor interface {
	Execute(ctx context.Context, id string) error
}

type updateProcessor interface {
	Execute(ctx context.Context, in DTO.Part) (DTO.Part, error)
}

type Controller struct {
	createProcessor createProcessor
	getProcessor    getProcessor
	deleteProcessor deleteProcessor
	updateProcessor updateProcessor
}

func New(
	createProcessor createProcessor,
	getProcessor getProcessor,
	deleteProcessor deleteProcessor,
	updateProcessor updateProcessor) *Controller {
	return &Controller{
		createProcessor: createProcessor,
		getProcessor:    getProcessor,
		deleteProcessor: deleteProcessor,
		updateProcessor: updateProcessor,
	}
}

func (c *Controller) CreateParts(w http.ResponseWriter, r *http.Request) {
	var in DTO.Part
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := c.createProcessor.Execute(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

func (c *Controller) GetParts(w http.ResponseWriter, r *http.Request) {
	parts, err := c.getProcessor.Execute(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, parts)
}

func (c *Controller) UpdateParts(w http.ResponseWriter, r *http.Request) {
	var in DTO.Part
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	in.Id = r.PathValue("id")

	updated, err := c.updateProcessor.Execute(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

func (c *Controller) DeleteParts(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := c.deleteProcessor.Execute(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
