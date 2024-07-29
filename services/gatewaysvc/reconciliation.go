package gatewaysvc

import (
	"context"
	"fmt"

	"canarails.dev/query"
)

// 全量同步
func Reconciliation(ctx context.Context) error {
	apps, err := query.App.WithContext(ctx).
		Preload(query.App.AppVariants).
		Find()
	if err != nil {
		return fmt.Errorf("find apps error: %w", err)
	}

	err = Apply(ctx, apps)
	if err != nil {
		return fmt.Errorf("apply gateway error: %w", err)
	}

	err = Prune(ctx, apps)
	if err != nil {
		return fmt.Errorf("prune gateway error: %w", err)
	}

	return nil
}
