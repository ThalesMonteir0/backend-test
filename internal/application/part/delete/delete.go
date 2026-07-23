package delete

import "context"

type service interface {
	DeletePart(ctx context.Context, id string) error
}
