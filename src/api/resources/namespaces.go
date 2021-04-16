package resources

import (
	"context"
	"net/http"

	c "infra-gw/src/cont"
	"infra-gw/src/util"
	"infra-gw/src/types"
	"infra-gw/src/api/helper"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
)

func ListNamespaces(ctx context.Context, appCtx *c.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := util.LoggerWithRequestID(appCtx.Log, r)

		log.Debug("list namespaces")

		namespaces, err := appCtx.K8s.Clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := make([]types.Namespaces, 0, len(namespaces.Items))

		for _, i := range namespaces.Items {
			result = append(result, types.Namespaces {
				Name: i.Name,
			})
		}

		resp := types.ListNamespacesResponse {
			Namespaces: result,
		}

		if err := helper.WriteJSONResponse(w, resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func CreateNamespaces(ctx context.Context, appCtx *c.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := util.LoggerWithRequestID(appCtx.Log, r)

		var req types.CreateNamespacesRequest
		err := helper.ParseBody(r.Body, &req)
		if err != nil {
			log.Errorw("Could not parse request body", "err", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Debugw("creating new namespace", "name", req.Name)
		bg := context.Background()

		ns := &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{Name: req.Name},
		}

		ns, err = appCtx.K8s.Clientset.CoreV1().Namespaces().Create(bg, ns, metav1.CreateOptions{})
		if err != nil {
			log.Errorw("Could not create namespace", "err", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Debugw("namespace was created", "name", req.Name)

		resp := types.CreateNamespacesResponse {
			Message: "OK",
			Status: http.StatusOK,
		}

		err = helper.WriteJSONResponse(w, resp)
		if err != nil {
			log.Errorw("Could not writeJSONResponse", "err", err.Error())
		}
	}
}

func DeleteNamespaces(ctx context.Context, appCtx *c.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := util.LoggerWithRequestID(appCtx.Log, r)

		var req types.DeleteNamespacesRequest
		err := helper.ParseBody(r.Body, &req)
		if err != nil {
			log.Errorw("Could not parse requst body", "err", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Debugw("Delete namespaces", "name", req.Name)
		bg := context.Background()

		err = appCtx.K8s.Clientset.CoreV1().Namespaces().Delete(bg, req.Name, metav1.DeleteOptions{})
		if err != nil {
			log.Errorw("Could not delete namespace", "name", req.Name)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Debugw("Namespace was deleted", "name", req.Name)

		resp := types.DeleteNamespacesResponse{
			Status: http.StatusOK,
			Message: "OK",
		}

		err = helper.WriteJSONResponse(w, resp)
		if err != nil {
			log.Errorw("Could not write WriteJSONResponse", "err", err.Error())
		}
	}
}
