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

func ListIngresses(ctx context.Context, appCtx *c.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := util.LoggerWithRequestID(appCtx.Log, r)

		log.Debug("list ingresses")

		ns, ok := r.URL.Query()["ns"]
		if !ok {
			ns = []string{""}
		}

		ingresses, err := appCtx.K8s.Clientset.ExtensionsV1beta1().Ingresses(ns[0]).List(ctx, metav1.ListOptions{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		result := make([]types.Ingresses, 0, len(ingresses.Items))

		for _, i := range ingresses.Items {
			result = append(result, types.Ingresses {
				Name: i.Name,
				Namespace: i.Namespace,
			})
		}

		resp := types.ListIngressesResponse {
			Ingresses: result,
		}

		if err := helper.WriteJSONResponse(w, resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}