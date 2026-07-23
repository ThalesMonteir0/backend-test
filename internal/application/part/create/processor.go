package create

import (
	"context"

	"github.com/ThalesMonteir0/backend-test/internal/DTO"
	domainpart "github.com/ThalesMonteir0/backend-test/internal/domain/part"
)

type Processor struct {
	service service
}

func New(service service) *Processor {
	return &Processor{service: service}
}

func (p *Processor) Execute(ctx context.Context, in DTO.Part) (DTO.Part, error) {
	created, err := p.service.CreatePart(ctx, domainpart.Part{
		ID:                in.Id,
		Name:              in.Name,
		Category:          in.Category,
		CurrentStock:      in.CurrentStock,
		MinimumStock:      in.MinimumStock,
		AverageDailySales: in.AverageDailySales,
		LeadTimeDays:      in.LeadTimeDays,
		CriticalityLevel:  in.CriticalityLevel,
		UnitCoast:         in.UnitCost,
	})
	if err != nil {
		return DTO.Part{}, err
	}

	return DTO.Part{
		Id:                created.ID,
		Name:              created.Name,
		Category:          created.Category,
		CurrentStock:      created.CurrentStock,
		MinimumStock:      created.MinimumStock,
		AverageDailySales: created.AverageDailySales,
		LeadTimeDays:      created.LeadTimeDays,
		CriticalityLevel:  created.CriticalityLevel,
		UnitCost:          created.UnitCoast,
	}, nil
}
