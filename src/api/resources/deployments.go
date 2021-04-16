package resources

import (
	"context"
	"net/http"

	c "infra-gw/src/cont"
	"infra-gw/src/util"
	"infra-gw/src/types"
	"infra-gw/src/api/helper"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListDeployments(ctx context.Context, appCtx *c.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := util.LoggerWithRequestID(appCtx.Log, r)

		log.Debug("list deployments")

		ns, ok := r.URL.Query()["ns"]
		if !ok {
			ns = []string{""}
		}

		deployments, err := appCtx.K8s.Clientset.ExtensionsV1beta1().Deployments(ns[0]).List(ctx, metav1.ListOptions{})
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
