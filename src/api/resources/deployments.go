package resources

import (
	"context"
	"net/http"

	c "infra-gw/src/cont"
	"infra-gw/src/util"
	"infra-gw/src/types"
	"infra-gw/src/api/helper"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func ListDeployments(ctx context.Context, appCtx *c.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := util.LoggerWithRequestID(appCtx.Log, r)

		log.Debug("list deployments")

		ns, ok := r.URL.Query()["ns"]
		if !ok {
			ns = []string{""}
		}

		deployments, err := appCtx.K8s.Clientset.AppsV1().Deployments(ns[0]).List(ctx, metav1.ListOptions{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := make([]types.Deployments, 0, len(deployments.Items))

		for _, i := range deployments.Items {
			result = append(result, types.Deployments {
				Name: i.Name,
				Namespace: i.Namespace,
			})
		}

		resp := types.ListDeploymentsResponse {
			Deployments: result,
		}

		if err := helper.WriteJSONResponse(w, resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func CreateDeployments(
		deploymentName string, 
		namespaceName string,
		envVars []corev1.EnvVar,
		resourceType string,
		imageName string,
		containerPort int,
		appCtx *c.AppContext) error {

	bg := context.Background()
	
	containerPort32 := int32(containerPort)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: util.Int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": resourceType,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": resourceType,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
							{
							Name: deploymentName,
							Image: imageName,
							Ports: []corev1.ContainerPort{
								{
									Name: deploymentName,
									Protocol: corev1.ProtocolTCP,
									ContainerPort: containerPort32,
								},
							},
							Env: envVars,
						},
					},
				},
			},
		},
	}

	_, err := appCtx.K8s.Clientset.AppsV1().Deployments(namespaceName).Create(bg, deployment, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}
