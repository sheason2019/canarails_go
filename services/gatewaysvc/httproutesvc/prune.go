package httproutesvc

import (
	"context"
	"fmt"
	"strconv"

	"canarails.dev/services/gatewaysvc/clientsvc"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func PruneHttpRoute(
	ctx context.Context,
	appIdMap map[uint]bool,
	appVarIdMap map[uint]bool,
) error {
	dynClient := clientsvc.NewDynamic()

	res := schema.GroupVersionResource{
		Group:    "gateway.networking.k8s.io",
		Version:  "v1",
		Resource: "httproutes",
	}

	httproutes, err := dynClient.Resource(res).
		Namespace("canarails").
		List(ctx, v1.ListOptions{})
	if err != nil {
		return fmt.Errorf("list http routes error: %w", err)
	}

	for _, item := range httproutes.Items {
		labels := item.GetLabels()

		if labels["isCanarailsResource"] != "true" {
			continue
		}

		appId, err := strconv.Atoi(labels["appId"])
		if err != nil {
			return fmt.Errorf("parse appId error: %w", err)
		}

		if !appIdMap[uint(appId)] {
			err = dynClient.Resource(res).
				Namespace("canarails").
				Delete(ctx, item.GetName(), v1.DeleteOptions{})
			if err != nil {
				return fmt.Errorf("delete http route error: %w", err)
			}
		}
	}

	return nil
}
