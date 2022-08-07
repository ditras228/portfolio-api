package desc

import (
	"context"
	"portfolio-api/graph/model"
)

type Repository interface {
	FindAll(ctx context.Context) ([]*model.GetDesc, error)
	UpdateDesc(ctx context.Context, input model.UpdateDescInput) (model.UpdateDescOutput, error)
	CreateDesc(ctx context.Context, input model.CreateDescInput) (model.CreateDescOutput, error)
	DeleteDesc(ctx context.Context, input model.DeleteDescInput) (model.DeleteDescOutput, error)
}
