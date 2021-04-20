package apps

import (
	"context"
	"net/http"

	c "infra-gw/src/cont"
	"infra-gw/src/util"
	"infra-gw/src/types/apps"
	"infra-gw/src/api/helper"
	"infra-gw/src/api/resources"

	corev1 "k8s.io/api/core/v1"
)

const redisType = "redis"
const redisImage = "redis:5.0.4"
const redisPort = 6379

func CreateRedis(ctx context.Context, appCtx *c.AppContext) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		log := util.LoggerWithRequestID(appCtx.Log, r)

		var req types.CreateRedisRequest
		err := helper.ParseBody(r.Body, &req)
		if err != nil {
			log.Errorw("Could not parse request body", "err", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Debugw("create new redis instance", "name", req.Name)

		envVars := []corev1.EnvVar{
			{Name: "MASTER", Value: "true"},
		}

		// Create deployment
		err = resources.CreateDeployments(req.Name, req.Namespace, envVars, redisType, redisImage, redisPort, appCtx)
		if err != nil {
			log.Errorw("Could not create deployment for redis", "err", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		// Create services
		err = resources.CreateServices(req.Name, req.Namespace, redisPort, appCtx)
		if err != nil {
			log.Errorw("Could not create service for redis", "err", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := types.CreateRedisResponse {
			Message: "OK",
			Status: http.StatusOK,
		}

		if err := helper.WriteJSONResponse(w, resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}