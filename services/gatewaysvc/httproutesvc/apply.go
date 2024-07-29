package httproutesvc

import (
	"context"
	"fmt"

	"canarails.dev/database/models"
	"canarails.dev/services/gatewaysvc/clientsvc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type json map[string]any

func createRules(app *models.App) []json {
	rules := make([]json, len(app.AppVariants))
	for i, appVar := range app.AppVariants {
		headers := make([]json, len(appVar.Matches))
		for j, match := range appVar.Matches {
			headers[j] = json{
				"type":  "Exact",
				"name":  match.Header,
				"value": match.Value,
			}
		}

		rules[i] = json{
			"matches": []json{
				{"headers": headers},
			},
			"filters": []json{
				{
					"type": "RequestHeaderModifier",
					"requestHeaderModifier": json{
						"add": []json{
							{
								"name":  "x-canarails-app-id",
								"value": fmt.Sprint(app.ID),
							},
							{
								"name":  "x-canarails-app-variant-id",
								"value": fmt.Sprint(appVar.ID),
							},
						},
					},
				},
				{
					"type": "ResponseHeaderModifier",
					"responseHeaderModifier": json{
						"add": []json{
							{
								"name":  "x-canarails-app-id",
								"value": fmt.Sprint(app.ID),
							},
							{
								"name":  "x-canarails-app-variant-id",
								"value": fmt.Sprint(appVar.ID),
							},
						},
					},
				},
			},
			"backendRefs": []json{
				{
					"name": fmt.Sprintf("service-%d", appVar.ID),
					"port": appVar.ExposePort,
				},
			},
		}
	}

	return rules
}

func ApplyHttpRoute(
	ctx context.Context,
	app *models.App,
) error {
	dynClient := clientsvc.NewDynamic()

	name := fmt.Sprintf("httproute-%d", app.ID)
	route := &unstructured.Unstructured{
		Object: json{
			"apiVersion": "gateway.networking.k8s.io/v1",
			"kind":       "HTTPRoute",
			"metadata": json{
				"name":      name,
				"namespace": "canarails",
				"labels": json{
					"isCanarailsResource": "true",
					"appId":               fmt.Sprint(app.ID),
				},
			},
			"spec": json{
				"parentRefs": []any{
					json{
						"name":      "canarails-gateway",
						"namespace": "canarails",
					},
				},
				"hostnames": app.Hostnames,
				"rules":     createRules(app),
			},
		},
	}
	res := schema.GroupVersionResource{
		Group:    "gateway.networking.k8s.io",
		Version:  "v1",
		Resource: "httproutes",
	}

	_, err := dynClient.Resource(res).
		Namespace("canarails").
		Apply(
			ctx,
			name,
			route,
			metav1.ApplyOptions{
				FieldManager: "application/apply-patch",
			},
		)
	return err
}
