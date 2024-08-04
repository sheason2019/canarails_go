package gatewaysvc

import (
	"context"
	"fmt"

	"canarails.dev/query"
	"gorm.io/gen/field"
)

// 全量同步
func Reconciliation(ctx context.Context, repo *query.Query) error {
	appVars, err := repo.AppVariant.WithContext(ctx).
		Join(repo.App, repo.AppVariant.AppID.EqCol(repo.App.ID)).
		Where(repo.AppVariant.Replicas.Gt(0)).
		Where(repo.AppVariant.ExposePort.Gt(0)).
		Where(repo.AppVariant.ImageName.NotLike("")).
		Where(field.Or(
			repo.AppVariant.Matches.Length().Gt(2),
			repo.AppVariant.ID.EqCol(repo.App.DefaultVariantID),
		)).
		Select(repo.AppVariant.ID).
		Find()
	if err != nil {
		return fmt.Errorf("query valid app variant error: %w", err)
	}

	appVarIds := make([]uint, len(appVars))
	for i, v := range appVars {
		appVarIds[i] = v.ID
	}

	apps, err := repo.App.WithContext(ctx).
		Join(repo.AppVariant, repo.AppVariant.AppID.EqCol(repo.App.ID)).
		Where(repo.App.Hostnames.Length().Gt(2)).
		Where(repo.AppVariant.ID.In(appVarIds...)).
		Preload(repo.App.AppVariants.On(
			repo.AppVariant.ID.In(appVarIds...),
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
