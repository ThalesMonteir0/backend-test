package priorieties

import (
	"context"
	domainpart "github.com/ThalesMonteir0/backend-test/internal/domain/part"
)

type service interface {
	PrioritizeParts(ctx context.Context) ([]domainpart.Part, error)
}
