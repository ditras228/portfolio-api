package info

import (
	"context"
	"portfolio-api/graph/model"
	"portfolio-api/pkg/utils"
)

func GetInfoForDTO(info model.GetInfo, name, exp model.GetTranslations, desc []*model.GetDesc, contacts model.Contacts) model.GetInfo {
	info.Name = &name
	info.Experience = utils.FormatHTML(exp)
	info.Desc = desc
	info.Contacts = &contacts
	return info
}

type Repository interface {
	FindOne(ctx context.Context) (model.GetInfo, error)
	UpdateInfo(ctx context.Context, input model.UpdateInfoInput) (model.GetInfo, error)
}
