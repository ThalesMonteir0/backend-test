package create

import (
	"context"

	domainpart "github.com/ThalesMonteir0/backend-test/internal/domain/part"
)

type service interface {
	CreatePart(ctx context.Context, p domainpart.Part) (domainpart.Part, error)
}
