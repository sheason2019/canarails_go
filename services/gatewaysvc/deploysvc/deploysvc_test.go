package deploysvc_test

import (
	"context"
	"testing"

	"canarails.dev/database"
	"canarails.dev/query"
	"canarails.dev/services/gatewaysvc/deploysvc"
)

func TestDeploy(t *testing.T) {
	ctx := context.Background()

	query.SetDefault(database.GetDb())
	app, err := query.App.WithContext(ctx).
		Preload(query.App.AppVariants).
		First()
	if err != nil {
		t.Error(err)
		return
	}

	for _, appVar := range app.AppVariants {
		err := deploysvc.ApplyDeployment(ctx, &appVar)
		if err != nil {
			t.Error(err)
			return
		}
	}
}
