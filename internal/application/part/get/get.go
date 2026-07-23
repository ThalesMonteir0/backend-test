package get

import (
	"context"

	domainpart "github.com/ThalesMonteir0/backend-test/internal/domain/part"
)

type service interface {
	GetParts(ctx context.Context) ([]domainpart.Part, error)
}
