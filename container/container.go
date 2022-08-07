package container

import (
	"context"
	"portfolio-api/infrastructure/postgresql"
	config "portfolio-api/internal"
	desc "portfolio-api/internal/desc/db"
	info "portfolio-api/internal/info/db"
	tag "portfolio-api/internal/tag/db"
	translation "portfolio-api/internal/translation/db"
	user "portfolio-api/internal/user/db"
	work "portfolio-api/internal/work/db"
)

var (
	cfg                 = config.GetConfig()
	client, _           = postgres.NewClient(context.TODO(), 3, cfg.Storage)
	TranslateRepository = translation.NewRepository(client)
	InfoRepository      = info.NewRepository(client, TranslateRepository, DescRepository)
	WorkRepository      = work.NewRepository(client, TranslateRepository, TagRepository)
	UserRepository      = user.NewRepository(client)
	TagRepository       = tag.NewRepository(client)
	DescRepository      = desc.NewRepository(client, TranslateRepository)
)
