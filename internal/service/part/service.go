package part

import (
	"context"

	domainpart "github.com/ThalesMonteir0/backend-test/internal/domain/part"
)

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreatePart(ctx context.Context, p domainpart.Part) (domainpart.Part, error) {
	return s.repo.Create(ctx, p)
}

func (s *Service) GetParts(ctx context.Context) ([]domainpart.Part, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) UpdatePart(ctx context.Context, p domainpart.Part) (domainpart.Part, error) {
	return s.repo.Update(ctx, p)
}

func (s *Service) DeletePart(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) PrioritizeParts(ctx context.Context) ([]domainpart.Part, error) {
	parts, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return domainpart.Prioritize(parts), nil
}
