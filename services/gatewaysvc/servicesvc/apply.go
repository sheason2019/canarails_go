package servicesvc

import (
	"context"
	"fmt"

	"canarails.dev/database/models"
	"canarails.dev/services/gatewaysvc/clientsvc"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
)

func createService(
	appVar *models.AppVariant,
) *corev1.ServiceApplyConfiguration {
	svc := corev1.Service(
		fmt.Sprintf("service-%d", appVar.ID),
		"canarails",
	)
	svc.WithLabels(map[string]string{
		"isCanarailsResource": "true",
		"appId":               fmt.Sprint(appVar.AppID),
		"appVariantId":        fmt.Sprint(appVar.ID),
	})
	svc.WithSpec(&corev1.ServiceSpecApplyConfiguration{
		Selector: map[string]string{
			"appVariantId": fmt.Sprint(appVar.ID),
		},
		Ports: []corev1.ServicePortApplyConfiguration{
			*corev1.ServicePort().
				WithName("main-expose").
				WithProtocol(v1.ProtocolTCP).
				WithPort(int32(appVar.ExposePort)).
				WithTargetPort(intstr.FromInt(int(appVar.ExposePort))),
		},
	})

	return svc
}

func ApplyService(
	ctx context.Context,
	appVar *models.AppVariant,
) error {
	client := clientsvc.New()

	_, err := client.CoreV1().
		Services("canarails").
		Apply(
			ctx,
			createService(appVar),
			metav1.ApplyOptions{
				FieldManager: "application/apply-patch",
			},
		)

	return err
}
