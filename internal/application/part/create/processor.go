package create

import (
	"context"
	"github.com/ThalesMonteir0/backend-test/internal/DTO"
	domainpart "github.com/ThalesMonteir0/backend-test/internal/domain/part"
	"go.uber.org/zap"
)

type Processor struct {
	service service
	logger  *zap.Logger
}

func New(service service, log *zap.Logger) *Processor {
	return &Processor{
		service: service,
		logger:  log,
	}
}

func (p *Processor) Execute(ctx context.Context, in DTO.Part) (DTO.Part, error) {
	if ctx.Err() != nil {
		p.logger.Warn("context is canceled", zap.Error(ctx.Err()))
		return DTO.Part{}, ctx.Err()
	}

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
		p.logger.Warn("create part failed", zap.Error(err))
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
