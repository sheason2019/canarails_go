package gatewaysvc_test

import (
	"context"
	"testing"

	"canarails.dev/database"
	"canarails.dev/query"
	"canarails.dev/services/gatewaysvc"
)

func TestReconciliation(t *testing.T) {
	ctx := context.Background()
	query.SetDefault(database.GetDb())

	err := gatewaysvc.Reconciliation(ctx, query.Q)
	if err != nil {
		t.Error(err)
	}
}
