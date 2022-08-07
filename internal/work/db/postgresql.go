package work

import (
	"context"
	"github.com/ztrue/tracerr"
	"portfolio-api/enitity"
	"portfolio-api/graph/model"
	"portfolio-api/infrastructure/postgresql"
	"portfolio-api/internal/tag"
	"portfolio-api/internal/translation"
	"portfolio-api/internal/work"
	"portfolio-api/middlewares/keys"
)

type repository struct {
	client          postgres.Client
	translationRepo translation.Repository
	tagRepo         tag.Repository
}

func (r *repository) FindAll(ctx context.Context) ([]*model.GetWork, error) {
	qWork := `

		SELECT 
			ID, name, description,
			github, demo

		FROM public.work 

		ORDER
			BY ID

		`

	workRows, err := r.client.Query(ctx, qWork)
	if err != nil {
		return nil, err
	}

	var nameTranslation model.GetTranslations
	var descTranslation model.GetTranslations
	var tags []*model.GetTag

	works := make([]*model.GetWork, 0)
	for workRows.Next() {
		var wrk model.GetWork

		err = workRows.Scan(&wrk.ID, &nameTranslation.Field, &descTranslation.Field,
			&wrk.Github, &wrk.Demo)
		if err != nil {
			return nil, err
		}
		tags, err = r.tagRepo.FindOne(ctx, wrk.ID)
		if err != nil {
			return nil, err
		}

		nameTranslation, err = r.translationRepo.FindOne(ctx, wrk.ID, enitity.WorkTitle, nameTranslation.Field)
		if err != nil {
			return nil, err
		}

		descTranslation, err = r.translationRepo.FindOne(ctx, wrk.ID, enitity.WorkFunctional, descTranslation.Field)
		if err != nil {
			return nil, err
		}

		wrk = work.GetWorkToDTO(wrk, nameTranslation, descTranslation, tags)
		works = append(works, &wrk)
	}
	if err != nil {
		return nil, err
	}

	return works, nil
}

func (r *repository) CreateWork(ctx context.Context, input model.CreateWorkInput) (model.GetWork, error) {
	qUpdWork := `

				INSERT INTO
					public.work
					(name, description, github, demo, figma)

				VALUES
					($1,$2,$3,$4,$5)
			
				RETURNING
					ID, name, description, github, demo

			   `

	var wrk model.GetWork
	var nameTranslation model.GetTranslations
	var descTranslation model.GetTranslations
	var tags []*model.GetTag
	var locale = keys.LocaleForContext(ctx) - 1

	err := r.client.
		QueryRow(ctx, qUpdWork,
			input.Name.Translations[locale].Field,
			input.Description.Translations[locale].Field,
			input.Github, input.Demo, input.Figma).
		Scan(&wrk.ID, &nameTranslation.Field, &descTranslation.Field, &wrk.Github, &wrk.Demo)
	if err != nil {
		return model.GetWork{}, err
	}
	nameTranslation, err = r.translationRepo.Create(ctx, input.Name, wrk.ID, enitity.WorkTitle, nameTranslation.Field)
	if err != nil {
		return model.GetWork{}, err
	}

	descTranslation, err = r.translationRepo.Create(ctx, input.Description, wrk.ID, enitity.WorkFunctional, descTranslation.Field)
	if err != nil {
		return model.GetWork{}, err
	}
	tags, err = r.tagRepo.Create(ctx, wrk.ID, input.Tags)
	if err != nil {
		return model.GetWork{}, err
	}

	wrk = work.GetWorkToDTO(wrk, nameTranslation, descTranslation, tags)

	return wrk, nil
}
func (r *repository) UpdateWork(ctx context.Context, input model.UpdateWorkInput) (model.UpdateWorkOutput, error) {

	qUpdWork := `

		UPDATE
			public.work

		SET
			name = $2, description = $3,
			github = $4, demo = $5, figma = $6

		WHERE 
			ID = $1

		RETURNING 
			ID, name, description, github, demo, figma

		`

	var wrk model.GetWork
	var nameTranslation model.GetTranslations
	var descTranslation model.GetTranslations
	var tags []*model.GetTag

	var locale = keys.LocaleForContext(ctx) - 1
	err := r.client.
		QueryRow(ctx, qUpdWork,
			input.ID, input.Name.Translations[locale].Field,
			input.Description.Translations[locale].Field, input.Github,
			input.Demo, input.Figma).
		Scan(&wrk.ID, &nameTranslation.Field, &descTranslation.Field, &wrk.Github, &wrk.Demo, &wrk.Figma)
	if err != nil {
		return model.NotFoundError{Message: "Работа не найдена", ID: input.ID}, nil
	}

	nameTranslation, err = r.translationRepo.Update(ctx, input.Name, wrk.ID, enitity.WorkTitle, nameTranslation.Field)
	if err != nil {
		return nil, err
	}

	descTranslation, err = r.translationRepo.Update(ctx, input.Description, wrk.ID, enitity.WorkFunctional, descTranslation.Field)
	if err != nil {
		return nil, err
	}
	tags, err = r.tagRepo.UpdateOne(ctx, input.ID, input.Tags)
	if err != nil {
		return nil, err
	}
	wrk = work.GetWorkToDTO(wrk, nameTranslation, descTranslation, tags)
	return wrk, nil
}

func (r *repository) DeleteWork(ctx context.Context, input model.DeleteWorkInput) (model.DeleteWorkOutput, error) {
	qWork := `

			DELETE FROM 
				public.work

			WHERE 
				ID = $1

			RETURNING 
				ID

		`

	var res model.DeleteWorkResult
	_, err := r.tagRepo.Delete(ctx, input.ID)
	if err != nil {
		return nil, tracerr.Errorf("Не удалось удалить теги: %s", err)
	}

	err = r.client.QueryRow(ctx, qWork, input.ID).Scan(&res.ID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewRepository(client postgres.Client, translationRepo translation.Repository, tagRepo tag.Repository) work.Repository {
	return &repository{
		client:          client,
		translationRepo: translationRepo,
		tagRepo:         tagRepo,
	}
}
