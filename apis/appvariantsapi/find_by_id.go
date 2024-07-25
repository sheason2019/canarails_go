package appvariantsapi

import (
	"context"
	"fmt"

	"canarails.dev/apis/genapi"
	"canarails.dev/query"
)

func (Impl) AppVariantsFindById(
	ctx context.Context,
	request genapi.AppVariantsFindByIdRequestObject,
) (genapi.AppVariantsFindByIdResponseObject, error) {
	record, err := query.AppVariant.WithContext(ctx).
		Where(query.AppVariant.ID.Eq(uint(request.Id))).
		First()
	if err != nil {
		return nil, fmt.Errorf("find app variant error: %w", err)
	}

	matches := make([]genapi.AppVariantMatch, len(record.Matches))
	for i, v := range record.Matches {
		matches[i] = genapi.AppVariantMatch{
			Header: v.Header,
			Value:  v.Value,
		}
	}

	appVariant := genapi.AppVariant{
		Id:          int32(record.ID),
		Title:       record.Title,
		Description: record.Description,
		ExposePort:  int32(record.ExposePort),
		Matches:     matches,
		AppId:       int32(record.AppID),
	}

	return genapi.AppVariantsFindById200JSONResponse(appVariant), nil
}
