package deploysvc

import (
	"context"
	"fmt"
	"strconv"

	"canarails.dev/services/gatewaysvc/clientsvc"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func PruneDeployment(
	ctx context.Context,
	appIdMap map[uint]bool,
	appVarIdMap map[uint]bool,
) error {
	client := clientsvc.New()
	list, err := client.AppsV1().
		Deployments("canarails").
		List(ctx, v1.ListOptions{})
	if err != nil {
		return fmt.Errorf("list deployment error: %w", err)
	}

	for _, item := range list.Items {
		labels := item.GetLabels()
		if labels["isCanarailsResource"] != "true" {
			continue
		}

		appId, err := strconv.Atoi(labels["appId"])
		if err != nil {
			return fmt.Errorf("parse appId error: %w", err)
		}

		appVarId, err := strconv.Atoi(labels["appVariantId"])
		if err != nil {
			return fmt.Errorf("parse appVariantId error: %w", err)
		}

		if !appIdMap[uint(appId)] || !appVarIdMap[uint(appVarId)] {
			err = client.AppsV1().
				Deployments("canarails").
				Delete(
					ctx,
					item.GetName(),
					v1.DeleteOptions{},
				)
			if err != nil {
				return fmt.Errorf("delete http route error: %w", err)
			}
		}
	}

	return nil
}
