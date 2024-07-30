package gatewaysvc

import (
	"context"
	"fmt"

	"canarails.dev/query"
	"gorm.io/gen/field"
)

// 全量同步
func Reconciliation(ctx context.Context) error {
	apps, err := query.App.WithContext(ctx).
		Join(query.AppVariant, query.AppVariant.AppID.EqCol(query.App.ID)).
		Where(query.AppVariant.Replicas.Gt(0)).
		Where(query.AppVariant.ImageName.NotLike("")).
		Where(field.Or(
			query.AppVariant.Matches.Length().Gt(2),
			query.AppVariant.ID.EqCol(query.App.DefaultVariantID),
		)).
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
