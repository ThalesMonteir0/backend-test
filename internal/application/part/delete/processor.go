package delete

import "context"

type Processor struct {
	service service
}

func New(service service) *Processor {
	return &Processor{service: service}
}

func (p *Processor) Execute(ctx context.Context, id string) error {
	return p.service.DeletePart(ctx, id)
}
