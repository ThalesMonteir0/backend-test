package update

import (
	"context"
	"go.uber.org/zap"

	"github.com/ThalesMonteir0/backend-test/internal/DTO"
	domainpart "github.com/ThalesMonteir0/backend-test/internal/domain/part"
)

type Processor struct {
	service service
	logger  *zap.Logger
}

func New(service service, log *zap.Logger) *Processor {
	return &Processor{service: service, logger: log}
}

func (p *Processor) Execute(ctx context.Context, in DTO.Part) (DTO.Part, error) {
	if ctx.Err() != nil {
		p.logger.Error("Context error", zap.Error(ctx.Err()))
		return DTO.Part{}, ctx.Err()
	}

	updated, err := p.service.UpdatePart(ctx, domainpart.Part{
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
		p.logger.Error("Error updating part", zap.Error(err))
		return DTO.Part{}, err
	}

	return DTO.Part{
		Id:                updated.ID,
		Name:              updated.Name,
		Category:          updated.Category,
		CurrentStock:      updated.CurrentStock,
		MinimumStock:      updated.MinimumStock,
		AverageDailySales: updated.AverageDailySales,
		LeadTimeDays:      updated.LeadTimeDays,
		CriticalityLevel:  updated.CriticalityLevel,
		UnitCost:          updated.UnitCoast,
	}, nil
}
