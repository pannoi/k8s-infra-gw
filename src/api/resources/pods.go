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

func ListPods(ctx context.Context, appCtx *c.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := util.LoggerWithRequestID(appCtx.Log, r)

		log.Debug("list pods")

		ns, ok := r.URL.Query()["ns"]
		if !ok {
			ns = []string{""}
		}

		pods, err := appCtx.K8s.Clientset.CoreV1().Pods(ns[0]).List(ctx, metav1.ListOptions{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := make([]types.Pods, 0, len(pods.Items))

		for _, i := range pods.Items {
			result = append(result, types.Pods {
				Name: i.Name,
				Namespace: i.Namespace,
			})
		}

		resp := types.ListPodsResponse {
			Pods: result,
		}

		if err := helper.WriteJSONResponse(w, resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
