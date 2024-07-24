package appsapi

import (
	"context"
	"fmt"

	"canarails.dev/apis/genapi"
	"canarails.dev/query"
)

func (Impl) AppsFindById(
	ctx context.Context,
	request genapi.AppsFindByIdRequestObject,
) (genapi.AppsFindByIdResponseObject, error) {
	record, err := query.App.WithContext(ctx).
		Where(query.App.ID.Eq(uint(request.Id))).
		First()
	if err != nil {
		return nil, fmt.Errorf("query app by id error: %w", err)
	}

	return genapi.AppsFindById200JSONResponse(genapi.App{
		Id:          int32(record.ID),
		Title:       record.Title,
		Description: record.Description,
		Hostnames:   record.Hostnames,
	}), nil
}
