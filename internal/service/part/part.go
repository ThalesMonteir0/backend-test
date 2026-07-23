package part

import (
	"context"

	domainpart "github.com/ThalesMonteir0/backend-test/internal/domain/part"
)

type Repository interface {
	Create(ctx context.Context, p domainpart.Part) (domainpart.Part, error)
	GetAll(ctx context.Context) ([]domainpart.Part, error)
	Update(ctx context.Context, p domainpart.Part) (domainpart.Part, error)
	Delete(ctx context.Context, id string) error
}
