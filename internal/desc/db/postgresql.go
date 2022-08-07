package desc

import (
	"context"
	"fmt"
	"github.com/ztrue/tracerr"
	"portfolio-api/enitity"
	"portfolio-api/graph/model"
	postgres "portfolio-api/infrastructure/postgresql"
	"portfolio-api/internal/desc"
	"portfolio-api/internal/translation"
	"portfolio-api/middlewares/keys"
	"portfolio-api/pkg/utils"
)

type repository struct {
	client          postgres.Client
	translationRepo translation.Repository
}

func (r *repository) FindAll(ctx context.Context) ([]*model.GetDesc, error) {
	q := `

			SELECT 
				ID, text, img

			FROM 
				public.desc

		 `

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, tracerr.Errorf("Не удалось найти все описания: %s", err)
	}
	var res []*model.GetDesc
	descs := make([]*model.GetDesc, 0)
	for rows.Next() {
		var dsc model.GetDesc
		var text model.GetTranslations
		err := rows.Scan(&dsc.ID, &text.Field, &dsc.Img)
		if err != nil {
			return nil, nil
		}

		dscTextTranslate, err := r.translationRepo.FindOne(ctx, dsc.ID, enitity.InfoDesc, text.Field)
		if err != nil {
			return nil, err
		}
		dsc.Text = &dscTextTranslate
		descs = append(descs, &dsc)
		fmt.Println(dsc)
	}
	res = descs
	return res, nil
}

func (r *repository) UpdateDesc(ctx context.Context, input model.UpdateDescInput) (model.UpdateDescOutput, error) {
	qImg := `

				SELECT 
					img
				
				FROM 
					public.desc

				WHERE 
					ID = $1

				`
	var oldLink string
	var newLink string

	err := r.client.
		QueryRow(ctx, qImg, input.ID).
		Scan(&oldLink)
	if err != nil {
		return model.NotFoundError{Message: "Описание не найдено", ID: input.ID}, nil
	}

	var dsc model.GetDesc
	var text model.GetTranslations

	qDesc := `

				UPDATE
					public.desc
				
				SET 
					text = $2, img = $3

				WHERE 
					ID = $1

				RETURNING
					ID, text, img

			 `

	newLink, err = utils.ReplaceImage(oldLink, input.Img)
	if err != nil {
		return nil, tracerr.Errorf("Не удалось заменить изображение: %s", err)
	}

	var locale = keys.LocaleForContext(ctx) - 1
	err = r.client.
		QueryRow(ctx, qDesc, input.ID, input.Text.Translations[locale].Field, newLink).
		Scan(&dsc.ID, &text.Field, &dsc.Img)
	if err != nil {
		return nil, tracerr.Errorf("Не удалось обновить описание: %s", err)
	}

	text, err = r.translationRepo.Update(ctx, input.Text, dsc.ID, enitity.InfoDesc, text.Field)
	if err != nil {
		return nil, err
	}

	dsc.Text = &text

	return dsc, nil
}

func (r *repository) CreateDesc(ctx context.Context, input model.CreateDescInput) (model.CreateDescOutput, error) {
	qDesc := `

				INSERT INTO 
					public.desc (text, img)

				VALUES
					($1, $2)

				RETURNING
					ID, text, img

			 `

	var dsc model.GetDesc
	var text model.Translation

	link, err := utils.SaveImage(input.Img)
	if err != nil {
		return nil, tracerr.Errorf("Не удалось сохранить картинку: %s", err)
	}
	var locale = keys.LocaleForContext(ctx) - 1
	err = r.client.
		QueryRow(ctx, qDesc, &input.Text.Translations[locale].Field, link).
		Scan(&dsc.ID, &text.Field, &dsc.Img)

	if err != nil {
		return nil, tracerr.Errorf("Не удалось создать описание: %s", err)
	}

	textUpd, err := r.translationRepo.Update(ctx, input.Text, dsc.ID, enitity.InfoDesc, text.Field)
	if err != nil {
		return nil, err
	}

	dsc.Text = &textUpd

	return dsc, nil
}

func (r *repository) DeleteDesc(ctx context.Context, input model.DeleteDescInput) (model.DeleteDescOutput, error) {
	qDesc := `

			DELETE FROM 
				public.desc
	
			WHERE 
				ID = $1

			RETURNING 
				ID

			`

	var res model.DeleteDescResult

	err := r.client.
		QueryRow(ctx, qDesc, input.ID).
		Scan(&res.ID)

	if err != nil {
		return model.NotFoundError{Message: "Описание не найдено", ID: input.ID}, nil
	}
	return res, nil
}
func NewRepository(client postgres.Client, translationRepo translation.Repository) desc.Repository {
	return &repository{
		client:          client,
		translationRepo: translationRepo,
	}
}
