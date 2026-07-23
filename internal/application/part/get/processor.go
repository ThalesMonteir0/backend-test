package get

import (
	"context"

	"github.com/ThalesMonteir0/backend-test/internal/DTO"
)

type Processor struct {
	service service
}

func New(service service) *Processor {
	return &Processor{service: service}
}

func (p *Processor) Execute(ctx context.Context) ([]DTO.Part, error) {
	parts, err := p.service.GetParts(ctx)
	if err != nil {
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
