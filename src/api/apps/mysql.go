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

const mysqlType = "mysql"
const mysqlImage = "mysql:5.6"
const mysqlPort = 3306

func CreateMySQL(ctx context.Context, appCtx *c.AppContext) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		log := util.LoggerWithRequestID(appCtx.Log, r)

		var req types.CreateMySQLRequest
		err := helper.ParseBody(r.Body, &req)
		if err != nil {
			log.Errorw("Could not parse request body", "err", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Debugw("create new mysql instance", "name", req.Namespace)

		// Create Secrets
		dbSecrets := map[string][]byte{
			"ROOT_PASSWORD": util.MustGenerateRandomAsByte(16),
		}

		err = resources.CreateSecrets(req.Name, req.Namespace, dbSecrets, appCtx)
		if err != nil {
			log.Errorw("could not create db secret", "err", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		envVars := []corev1.EnvVar{
			{Name: "MYSQL_ROOT_PASSWORD", ValueFrom: util.SecretKeyRef(req.Name, "ROOT_PASSWORD")},
		}

		// Create Deployment
		err = resources.CreateDeployments(req.Name, req.Namespace, envVars, mysqlType, mysqlImage, mysqlPort, appCtx)
		if err != nil {
			log.Errorw("Could not create deployment for mysql", "err", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		// Create Service
		err = resources.CreateServices(req.Name, req.Namespace, mysqlPort, appCtx)
		if err != nil {
			log.Errorw("Could not create service for mysql", "err", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Create response
		resp := types.CreateMySQLResponse {
			Message: "OK",
			Status: http.StatusOK,
			DatabaseUsername: "root",
			DatabasePassword: dbSecrets,
		}

		if err := helper.WriteJSONResponse(w, resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}