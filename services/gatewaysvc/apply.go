package gatewaysvc

import (
	"context"
	"fmt"

	"canarails.dev/database/models"
	"canarails.dev/services/gatewaysvc/deploysvc"
	"canarails.dev/services/gatewaysvc/httproutesvc"
	"canarails.dev/services/gatewaysvc/servicesvc"
)

func Apply(
	ctx context.Context,
	apps []*models.App,
) error {
	for _, app := range apps {
		err := httproutesvc.ApplyHttpRoute(ctx, app)
		if err != nil {
			return fmt.Errorf("apply http route error: %w", err)
		}

		for _, appVar := range app.AppVariants {
			err = servicesvc.ApplyService(ctx, &appVar)
			if err != nil {
				return fmt.Errorf("apply service error: %w", err)
			}

			err = deploysvc.ApplyDeployment(ctx, &appVar)
			if err != nil {
				return fmt.Errorf("apply deployment error: %w", err)
			}
		}
	}

	return nil
}
