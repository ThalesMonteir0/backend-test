package delete

import (
	"context"
	"go.uber.org/zap"
)

type Processor struct {
	service service
	logger  *zap.Logger
}

func New(service service, log *zap.Logger) *Processor {
	return &Processor{service: service, logger: log}
}

func (p *Processor) Execute(ctx context.Context, id string) error {
	if ctx.Err() != nil {
		p.logger.Warn("context canceled", zap.Error(ctx.Err()))
		return ctx.Err()
	}

	return p.service.DeletePart(ctx, id)
}
