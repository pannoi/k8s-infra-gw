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

func ListServices(ctx context.Context, appCtx *c.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := util.LoggerWithRequestID(appCtx.Log, r)

		log.Debug("list services")

		ns, ok := r.URL.Query()["ns"]
		if !ok {
			ns = []string{""}
		}

		services, err := appCtx.K8s.Clientset.CoreV1().Services(ns[0]).List(ctx, metav1.ListOptions{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := make([]types.Services, 0, len(services.Items))

		for _, i := range services.Items {
			result = append(result, types.Services {
				Name: i.Name,
				Namespace: i.Namespace,
			})
		}
		
		resp := types.ListServicesResponse {
			Services: result,
		}

		if err := helper.WriteJSONResponse(w, resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
