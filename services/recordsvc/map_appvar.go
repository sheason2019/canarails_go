package recordsvc

import (
	"canarails.dev/apis/genapi"
	"canarails.dev/database/models"
)

func MapAppVar(data *genapi.OptionalAppVariant) *models.AppVariant {
	output := models.AppVariant{}

	if data.Id != nil {
		output.ID = uint(*data.Id)
	}
	if data.AppId != nil {
		output.AppID = uint(*data.AppId)
	}
	if data.Title != nil {
		output.Title = *data.Title
	}
	if data.Description != nil {
		output.Description = *data.Description
	}
	if data.ImageName != nil {
		output.ImageName = *data.ImageName
	}
	if data.ExposePort != nil {
		output.ExposePort = uint(*data.ExposePort)
	}
	if data.Replicas != nil {
		output.Replicas = uint(*data.Replicas)
	}
	if data.Matches != nil {
		matches := make([]models.AppVariantMatch, len(*data.Matches))
		for i, v := range *data.Matches {
			matches[i] = models.AppVariantMatch{
				Header: v.Header,
				Value:  v.Value,
			}
		}
		output.Matches = matches
	}

	return &output
}
