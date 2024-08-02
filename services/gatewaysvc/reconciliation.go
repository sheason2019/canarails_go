package gatewaysvc

import (
	"context"
	"fmt"

	"canarails.dev/query"
	"gorm.io/gen/field"
)

// 全量同步
func Reconciliation(ctx context.Context) error {
	appVars, err := query.AppVariant.WithContext(ctx).
		Join(query.App, query.AppVariant.AppID.EqCol(query.App.ID)).
		Where(query.AppVariant.Replicas.Gt(0)).
		Where(query.AppVariant.ImageName.NotLike("")).
		Where(field.Or(
			query.AppVariant.Matches.Length().Gt(2),
			query.AppVariant.ID.EqCol(query.App.DefaultVariantID),
		)).
		Select(query.AppVariant.ID).
		Find()
	if err != nil {
		return fmt.Errorf("query valid app variant error: %w", err)
	}

	appVarIds := make([]uint, len(appVars))
	for i, v := range appVars {
		appVarIds[i] = v.ID
	}

	apps, err := query.App.WithContext(ctx).
		Join(query.AppVariant, query.AppVariant.AppID.EqCol(query.App.ID)).
		Where(query.App.Hostnames.Length().Gt(2)).
		Where(query.AppVariant.ID.In(appVarIds...)).
		Preload(query.App.AppVariants.On(
			query.AppVariant.ID.In(appVarIds...),
		)).
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
