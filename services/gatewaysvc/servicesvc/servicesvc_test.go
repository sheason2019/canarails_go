package servicesvc_test

import (
	"context"
	"testing"

	"canarails.dev/database"
	"canarails.dev/query"
	"canarails.dev/services/gatewaysvc/servicesvc"
)

func TestApplyService(t *testing.T) {
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
		err := servicesvc.ApplyService(ctx, &appVar)
		if err != nil {
			t.Error(err)
			return
		}
	}
}
