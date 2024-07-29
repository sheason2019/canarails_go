package httproutesvc_test

import (
	"context"
	"testing"

	"canarails.dev/database"
	"canarails.dev/query"
	"canarails.dev/services/gatewaysvc/httproutesvc"
)

func TestApplyHttpRoute(t *testing.T) {
	ctx := context.Background()
	query.SetDefault(database.GetDb())

	app, err := query.App.WithContext(ctx).
		Preload(query.App.AppVariants).
		First()
	if err != nil {
		t.Error(err)
		return
	}

	err = httproutesvc.ApplyHttpRoute(ctx, app)
	if err != nil {
		t.Error(err)
		return
	}
}
