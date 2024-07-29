package deploysvc

import (
	"context"
	"fmt"

	"canarails.dev/database/models"
	"canarails.dev/services/gatewaysvc/clientsvc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1 "k8s.io/client-go/applyconfigurations/apps/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

func createDeployment(
	appVar *models.AppVariant,
) *appsv1.DeploymentApplyConfiguration {
	return appsv1.Deployment(
		fmt.Sprintf("deployment-%d", appVar.ID),
		"canarails",
	).WithLabels(
		map[string]string{
			"isCanarailsResource": "true",
			"appId":               fmt.Sprint(appVar.AppID),
			"appVariantId":        fmt.Sprint(appVar.ID),
		},
	).WithSpec(
		appsv1.DeploymentSpec().
			WithReplicas(int32(appVar.Replicas)).
			WithSelector(
				v1.LabelSelector().
					WithMatchLabels(map[string]string{
						"appId":        fmt.Sprint(appVar.AppID),
						"appVariantId": fmt.Sprint(appVar.ID),
					}),
			).
			WithTemplate(
				corev1.PodTemplateSpec().
					WithLabels(map[string]string{
						"appId":        fmt.Sprint(appVar.AppID),
						"appVariantId": fmt.Sprint(appVar.ID),
					}).
					WithSpec(
						corev1.PodSpec().WithContainers(
							corev1.Container().
								WithName(fmt.Sprintf("deployment-%d", appVar.ID)).
								WithImage(appVar.ImageName).
								WithPorts(
									corev1.ContainerPort().
										WithContainerPort(int32(appVar.ExposePort)),
								),
						),
					),
			),
	)
}

func ApplyDeployment(
	ctx context.Context,
	appVar *models.AppVariant,
) error {
	client := clientsvc.New()

	_, err := client.AppsV1().
		Deployments("canarails").
		Apply(
			ctx,
			createDeployment(appVar),
			metav1.ApplyOptions{
				FieldManager: "application/apply-patch",
			},
		)

	return err
}
