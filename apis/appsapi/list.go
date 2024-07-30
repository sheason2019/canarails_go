package appsapi

import (
	"context"

	"canarails.dev/apis/genapi"
	"canarails.dev/query"
)

func (Impl) AppsList(
	ctx context.Context,
	request genapi.AppsListRequestObject,
) (genapi.AppsListResponseObject, error) {
	records, err := query.App.WithContext(ctx).Find()
	if err != nil {
		return nil, err
	}

	apps := make([]genapi.App, len(records))
	for i, v := range records {
		apps[i] = genapi.App{
			Id:               int32(v.ID),
			Title:            v.Title,
			Description:      v.Description,
			Hostnames:        v.Hostnames,
			DefaultVariantId: int32(v.DefaultVariantID),
		}
	}

	return genapi.AppsList200JSONResponse(apps), nil
}
