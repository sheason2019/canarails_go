package versionsapi

import (
	"context"

	"canarails.dev/apis/genapi"
	"canarails.dev/services/aboutsvc"
)

func (Impl) VersionsGetVersion(
	ctx context.Context,
	request genapi.VersionsGetVersionRequestObject,
) (genapi.VersionsGetVersionResponseObject, error) {
	v := genapi.Version{
		GitHash:   aboutsvc.GitHash,
		BuildTime: aboutsvc.BuildTime,
	}

	return genapi.VersionsGetVersion200JSONResponse(v), nil
}
