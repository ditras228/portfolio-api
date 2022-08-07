package tag

import (
	"context"
	"portfolio-api/graph/model"
)

type Repository interface {
	FindAll(ctx context.Context) ([]*model.GetTag, error)
	FindOne(ctx context.Context, workID int) ([]*model.GetTag, error)
	UpdateOne(ctx context.Context, workId int, tags []int) ([]*model.GetTag, error)
	Create(ctx context.Context, workId int, tags []int) ([]*model.GetTag, error)
	Delete(ctx context.Context, workId int) ([]*int, error)
}
