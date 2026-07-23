package restock

import (
	"context"
	"github.com/ThalesMonteir0/backend-test/internal/DTO"
	"github.com/ThalesMonteir0/backend-test/internal/jsonx"
	"net/http"
)

type prioritiesRestockProcessor interface {
	Execute(ctx context.Context) (DTO.RestockResponse, error)
}
type Controller struct {
	prioritiesRestockProcessor prioritiesRestockProcessor
}

func New(prioritiesRestockProcessor prioritiesRestockProcessor) *Controller {
	return &Controller{
		prioritiesRestockProcessor: prioritiesRestockProcessor,
	}
}

func (c *Controller) Priorities(w http.ResponseWriter, r *http.Request) {
	parts, err := c.prioritiesRestockProcessor.Execute(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonx.WriteJSON(w, http.StatusOK, parts)
}
