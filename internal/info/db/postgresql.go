package info

import (
	"context"
	"github.com/ztrue/tracerr"
	"portfolio-api/enitity"
	"portfolio-api/graph/model"
	"portfolio-api/infrastructure/postgresql"
	"portfolio-api/internal/desc"
	"portfolio-api/internal/info"
	"portfolio-api/internal/translation"
	"portfolio-api/middlewares/keys"
	"portfolio-api/pkg/utils"
)

type repository struct {
	client          postgres.Client
	translationRepo translation.Repository
	descRepo        desc.Repository
}

func (r *repository) UpdateInfo(ctx context.Context, input model.UpdateInfoInput) (model.GetInfo, error) {
	qImg := `

				SELECT 
					img
				
				FROM 
					public.info

				WHERE 
					ID = 1

			`

	var oldLink string
	var newLink string

	err := r.client.
		QueryRow(ctx, qImg).
		Scan(&oldLink)

	q := `

		UPDATE
			public.info

		SET
			name = $1, job = $2,  experience = $3,
			telegramTitle = $4, telegramLink = $5, githubTitle = $6, githubLink = $7, img = $8

		WHERE 
			id = 1

		RETURNING 
			name, job,  experience,
			telegramTitle, telegramLink, githubTitle, githubLink

		`

	var inf model.GetInfo
	var con model.Contacts

	if err != nil {
		return model.GetInfo{}, err
	}
	newLink, err = utils.ReplaceImage(oldLink, input.Img)
	if err != nil {
		return model.GetInfo{}, err
	}

	var name model.GetTranslations
	var exp model.GetTranslations

	var locale = keys.LocaleForContext(ctx) - 1
	err = r.client.QueryRow(ctx, q,
		input.Name.Translations[locale].Field, input.Job,
		input.Experience.Translations[locale].Field, input.TelegramTitle,
		input.TelegramLink, input.GithubTitle, input.GithubLink, newLink).
		Scan(&name.Field, &inf.Job, &exp.Field, &con.TelegramTitle, &con.TelegramLink, &con.GithubTitle, &con.GithubLink)

	nameUpd, err := r.translationRepo.Update(ctx, input.Name, 1, enitity.InfoTitle, name.Field)
	if err != nil {
		return model.GetInfo{}, err
	}
	expUpd, err := r.translationRepo.Update(ctx, input.Experience, 1, enitity.InfoExperience, name.Field)
	if err != nil {
		return model.GetInfo{}, err
	}

	descs, err := r.descRepo.FindAll(ctx)
	if err != nil {
		return model.GetInfo{}, err
	}
	inf = info.GetInfoForDTO(inf, nameUpd, expUpd, descs, con)
	return inf, nil
}

func (r *repository) FindOne(ctx context.Context) (model.GetInfo, error) {
	q := `

		SELECT 
			name, job, 
			experience, telegramTitle,
			telegramLink, githubTitle,
			githubLink, img

		FROM 
			public.info 

		WHERE 
			ID = 1

		`

	var inf model.GetInfo
	var con model.Contacts

	var name model.GetTranslations
	var exp model.GetTranslations
	err := r.client.QueryRow(ctx, q).Scan(
		&name.Field, &inf.Job, &exp.Field,
		&con.TelegramTitle, &con.TelegramLink, &con.GithubTitle, &con.GithubLink, &inf.Img,
	)
	if err != nil {
		return model.GetInfo{}, err
	}

	descs, err := r.descRepo.FindAll(ctx)
	if err != nil {
		return model.GetInfo{}, tracerr.Errorf("Ошибка получения описаний: %s", err)
	}

	infoNameTranslate, err := r.translationRepo.FindOne(ctx, 1, enitity.InfoTitle, name.Field)
	if err != nil {
		return model.GetInfo{}, tracerr.Errorf("Ошибка получения перевода имени: %s", err)
	}

	infExperienceTranslate, err := r.translationRepo.FindOne(ctx, 1, enitity.InfoExperience, exp.Field)
	if err != nil {
		return model.GetInfo{}, tracerr.Errorf("Ошибка получения перевода опыта: %s", err)
	}

	inf = info.GetInfoForDTO(inf, infoNameTranslate, infExperienceTranslate, descs, con)

	return inf, nil
}
func NewRepository(client postgres.Client, translationRepo translation.Repository, descRepo desc.Repository) info.Repository {
	return &repository{
		client:          client,
		translationRepo: translationRepo,
		descRepo:        descRepo,
	}
}
