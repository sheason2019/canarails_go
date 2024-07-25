package appvariantsapi

import (
	"context"
	"fmt"

	"canarails.dev/apis/genapi"
	"canarails.dev/query"
)

func (Impl) AppVariantsList(
	ctx context.Context,
	request genapi.AppVariantsListRequestObject,
) (genapi.AppVariantsListResponseObject, error) {
	records, err := query.AppVariant.WithContext(ctx).
		Where(query.AppVariant.AppID.Eq(uint(request.Params.AppId))).
		Find()
	if err != nil {
		return nil, fmt.Errorf("find app variants error: %w", err)
	}

	appVariants := make([]genapi.AppVariant, len(records))
	for i, v := range records {
		matches := make([]genapi.AppVariantMatch, len(v.Matches))
		for matchIdx, match := range v.Matches {
			matches[matchIdx] = genapi.AppVariantMatch{
				Header: match.Header,
				Value:  match.Value,
			}
		}

		appVariants[i] = genapi.AppVariant{
			Id:          int32(v.ID),
			AppId:       int32(v.AppID),
			Title:       v.Title,
			Description: v.Description,
			Matches:     matches,
			ExposePort:  int32(v.ExposePort),
		}
	}

	return genapi.AppVariantsList200JSONResponse(appVariants), nil
}
