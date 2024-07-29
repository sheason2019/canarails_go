package servicesvc

import (
	"context"
	"fmt"
	"strconv"

	"canarails.dev/services/gatewaysvc/clientsvc"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func PruneService(
	ctx context.Context,
	appIdMap map[uint]bool,
	appVarIdMap map[uint]bool,
) error {
	client := clientsvc.New()
	list, err := client.CoreV1().
		Services("canarails").
		List(ctx, v1.ListOptions{})
	if err != nil {
		return fmt.Errorf("list service error: %w", err)
	}

	for _, svc := range list.Items {
		labels := svc.GetLabels()
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
			err = client.CoreV1().
				Services("canarails").
				Delete(
					ctx,
					svc.GetName(),
					v1.DeleteOptions{},
				)
			if err != nil {
				return fmt.Errorf("delete http route error: %w", err)
			}
		}
	}

	return nil
}
