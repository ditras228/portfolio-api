package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"portfolio-api/container"
	"portfolio-api/graph/generated"
	"portfolio-api/graph/model"
)

func (r *mutationResolver) Auth(ctx context.Context, input model.UserInput) (model.UserOutput, error) {
	user, err := container.UserRepository.Auth(ctx, input)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mutationResolver) UpdateInfo(ctx context.Context, input model.UpdateInfoInput) (*model.GetInfo, error) {
	upd, err := container.InfoRepository.UpdateInfo(ctx, input)
	if err != nil {
		return nil, err
	}
	return &upd, nil
}

func (r *mutationResolver) CreateWork(ctx context.Context, input model.CreateWorkInput) (*model.GetWork, error) {
	res, err := container.WorkRepository.CreateWork(ctx, input)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *mutationResolver) UpdateWork(ctx context.Context, input model.UpdateWorkInput) (model.UpdateWorkOutput, error) {
	res, err := container.WorkRepository.UpdateWork(ctx, input)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) DeleteWork(ctx context.Context, input model.DeleteWorkInput) (model.DeleteWorkOutput, error) {
	res, err := container.WorkRepository.DeleteWork(ctx, input)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) UpdateDesc(ctx context.Context, input model.UpdateDescInput) (model.UpdateDescOutput, error) {
	desc, err := container.DescRepository.UpdateDesc(ctx, input)
	if err != nil {
		return nil, err
	}
	return desc, nil
}

func (r *mutationResolver) CreateDesc(ctx context.Context, input model.CreateDescInput) (model.CreateDescOutput, error) {
	desc, err := container.DescRepository.CreateDesc(ctx, input)
	if err != nil {
		return nil, err
	}
	return desc, nil
}

func (r *mutationResolver) DeleteDesc(ctx context.Context, input model.DeleteDescInput) (model.DeleteDescOutput, error) {
	desc, err := container.DescRepository.DeleteDesc(ctx, input)
	if err != nil {
		return nil, err
	}
	return desc, nil
}

func (r *queryResolver) GetInfo(ctx context.Context) (*model.GetInfo, error) {
	one, err := container.InfoRepository.FindOne(ctx)
	if err != nil {
		return nil, err
	}
	//res, err := container.TranslateRepository.FindOne(ctx, 1)
	//if err != nil {
	//	return nil, err
	//}
	//one.Name = res.Translations[0].Field
	//fmt.Println(len(res.Translations))
	return &one, nil
}

func (r *queryResolver) GetWorks(ctx context.Context) ([]*model.GetWork, error) {
	works, err := container.WorkRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return works, nil
}

func (r *queryResolver) GetTags(ctx context.Context) ([]*model.GetTag, error) {
	tags, err := container.TagRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *queryResolver) GetDesc(ctx context.Context) ([]*model.GetDesc, error) {
	res, err := container.DescRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) GetOneUser(ctx context.Context, id int) (*model.User, error) {
	res, err := container.UserRepository.GetOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
