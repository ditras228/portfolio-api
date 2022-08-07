package translation

import (
	"context"
	"github.com/ztrue/tracerr"
	"portfolio-api/graph/model"
	"portfolio-api/infrastructure/postgresql"
	"portfolio-api/internal/translation"
	"portfolio-api/middlewares/keys"
	"regexp"
	"strconv"
	"strings"
)

type repository struct {
	client postgres.Client
}

func (r *repository) FindOne(ctx context.Context, translateId, entityId int, origValue string) (model.GetTranslations, error) {
	q := `

		SELECT 
			field, locale

		FROM 
			public.translation

		WHERE 
			translateID = $1 AND entityID = $2 

		`

	rows, err := r.client.Query(ctx, q, translateId, entityId)
	if err != nil {
		return model.GetTranslations{}, err
	}

	translations := make([]*model.Translation, 0)
	var newField string
	for rows.Next() {
		var trn model.Translation

		err := rows.Scan(&trn.Field, &trn.Locale)
		if err != nil {
			return model.GetTranslations{}, err
		}

		translations = append(translations, &trn)
		if trn.Locale == keys.LocaleForContext(ctx) {
			newField = trn.Field
		} else {
			newField = origValue
		}
	}

	// На тот случай, если нужно выводить все переводы, а орига в базе нет
	if len(translations) == 1 {
		var mockTranslation model.Translation
		mockTranslation.Locale = 2
		mockTranslation.Field = origValue
		translations = append(translations, &mockTranslation)
	}
	return model.GetTranslations{Field: newField, Translations: translations}, err
}
func (r *repository) Create(ctx context.Context, input *model.UpdateTranslationInput, translationId, entityId int, origValue string) (model.GetTranslations, error) {
	qAddTranslations := `

						INSERT INTO 
							public.translation (translateID, entityID, locale, field)
			
						VALUES
			
						`

	for i := 0; i < len(input.Translations); i++ {
		fieldRegex := regexp.MustCompile("\n")
		cleanTranslation := fieldRegex.ReplaceAllString(input.Translations[i].Field, "")
		var values []string
		values = append(values,
			strconv.Itoa(translationId),
			strconv.Itoa(entityId),
			strconv.Itoa(input.Translations[i].Locale),
			cleanTranslation)

		valuesStr := strings.Join(values, "', '")
		var qAddTranslationItem = "('" + valuesStr + "'),"
		qAddTranslations = qAddTranslations + qAddTranslationItem
	}

	qAddTranslations = qAddTranslations[0 : len(qAddTranslations)-1]

	rows, err := r.client.Query(ctx, qAddTranslations)
	if err != nil {
		return model.GetTranslations{}, err
	}

	translations := make([]*model.Translation, 0)
	for rows.Next() {
		var trn model.Translation

		err := rows.Scan(&trn.Field, &trn.Locale)
		if err != nil {
			return model.GetTranslations{}, err
		}

		translations = append(translations, &trn)
	}

	return model.GetTranslations{Field: origValue, Translations: translations}, err
}
func (r *repository) Delete(ctx context.Context, translationId, entityId int) (int, error) {
	qDelete := `
				DELETE FROM 
					public.translation

				WHERE 
					translateID = $1 AND entityID = $2

				RETURNING
					translateID
  			 `

	var id int
	err := r.client.QueryRow(ctx, qDelete, translationId, entityId).Scan(&id)
	if err != nil {
		return 0, tracerr.Errorf(`Не удалось удалить перевод: %s`, err)
	}
	return id, nil
}
func (r *repository) Update(ctx context.Context, input *model.UpdateTranslationInput, translationId, entityId int, origValue string) (model.GetTranslations, error) {
	_, err := r.Delete(ctx, translationId, entityId)
	if err != nil {
		return model.GetTranslations{}, err
	}
	translations, err := r.Create(ctx, input, translationId, entityId, origValue)
	if err != nil {
		return model.GetTranslations{}, err
	}
	return translations, nil

}
func NewRepository(client postgres.Client) translation.Repository {
	return &repository{
		client: client,
	}
}
