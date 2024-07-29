package gatewaysvc

import (
	"context"
	"fmt"

	"canarails.dev/database/models"
	"canarails.dev/services/gatewaysvc/deploysvc"
	"canarails.dev/services/gatewaysvc/httproutesvc"
	"canarails.dev/services/gatewaysvc/servicesvc"
)

// 裁剪不再使用的资源
func Prune(
	ctx context.Context,
	apps []*models.App,
) error {
	// 收集 ID
	appIdMap := map[uint]bool{}
	appVarIdMap := map[uint]bool{}

	for _, app := range apps {
		appIdMap[app.ID] = true
		for _, appVar := range app.AppVariants {
			appVarIdMap[appVar.ID] = true
		}
	}

	// prune httproute
	err := httproutesvc.PruneHttpRoute(ctx, appIdMap, appVarIdMap)
	if err != nil {
		return fmt.Errorf("prune http route error: %w", err)
	}

	// prune service
	err = servicesvc.PruneService(ctx, appIdMap, appVarIdMap)
	if err != nil {
		return fmt.Errorf("prune service error: %w", err)
	}

	// prune deployment
	err = deploysvc.PruneDeployment(ctx, appIdMap, appVarIdMap)
	if err != nil {
		return fmt.Errorf("prune deployment error: %w", err)
	}

	return nil
}
