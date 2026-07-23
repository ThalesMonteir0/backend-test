package priorieties

import (
	"context"
	"github.com/ThalesMonteir0/backend-test/internal/DTO"
	"go.uber.org/zap"
)

type Processor struct {
	service service
	logger  *zap.Logger
}

func NewProcessor(partService service, log *zap.Logger) *Processor {
	return &Processor{
		service: partService,
		logger:  log,
	}
}

func (p *Processor) Execute(ctx context.Context) (DTO.RestockResponse, error) {
	if ctx.Err() != nil {
		p.logger.Error("context canceled", zap.Error(ctx.Err()))
		return DTO.RestockResponse{}, ctx.Err()
	}

	parts, err := p.service.PrioritizeParts(ctx)
	if err != nil {
		p.logger.Error("prioritize parts error", zap.Error(err))
		return DTO.RestockResponse{}, err
	}

	priorities := make([]DTO.Priorities, 0, len(parts))
	for _, part := range parts {
		priorities = append(priorities, DTO.Priorities{
			Id:             part.ID,
			Name:           part.Name,
			CurrentStock:   part.CurrentStock,
			ProjectedStock: part.ProjectedStock(),
			MinimumStock:   part.MinimumStock,
			UrgencyStock:   part.UrgencyScore(),
		})
	}

	return DTO.RestockResponse{Priorities: priorities}, nil
}
