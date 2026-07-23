package get

import (
	"context"
	"go.uber.org/zap"

	"github.com/ThalesMonteir0/backend-test/internal/DTO"
)

type Processor struct {
	service service
	logger  *zap.Logger
}

func New(service service, log *zap.Logger) *Processor {
	return &Processor{service: service, logger: log}
}

func (p *Processor) Execute(ctx context.Context) ([]DTO.Part, error) {
	if ctx.Err() != nil {
		p.logger.Error("Context err", zap.Error(ctx.Err()))
		return []DTO.Part{}, ctx.Err()
	}

	parts, err := p.service.GetParts(ctx)
	if err != nil {
		p.logger.Error("Error getting parts", zap.Error(err))
		return nil, err
	}

	result := make([]DTO.Part, 0, len(parts))
	for _, part := range parts {
		result = append(result, DTO.Part{
			Id:                part.ID,
			Name:              part.Name,
			Category:          part.Category,
			CurrentStock:      part.CurrentStock,
			MinimumStock:      part.MinimumStock,
			AverageDailySales: part.AverageDailySales,
			LeadTimeDays:      part.LeadTimeDays,
			CriticalityLevel:  part.CriticalityLevel,
			UnitCost:          part.UnitCoast,
		})
	}

	return result, nil
}
