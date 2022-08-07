package translation

import (
	"context"
	"portfolio-api/graph/model"
)

type Repository interface {
	FindOne(ctx context.Context, translateId, entityId int, origValue string) (model.GetTranslations, error)
	Update(ctx context.Context, input *model.UpdateTranslationInput, translationId, entityId int, origValue string) (model.GetTranslations, error)
	Create(ctx context.Context, input *model.UpdateTranslationInput, translationId, entityId int, origValue string) (model.GetTranslations, error)
	Delete(ctx context.Context, translationId, entityId int) (int, error)
}
