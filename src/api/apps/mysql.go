package apps

import (
	"context"
	"net/http"

	c "infra-gw/src/cont"
	"infra-gw/src/util"
	"infra-gw/src/types/apps"
	"infra-gw/src/api/helper"
	"infra-gw/src/api/resources"
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

		// Create Deployment
		err = resources.CreateDeployments(req.Name, req.Namespace, mysqlType, mysqlImage, mysqlPort, appCtx)
		if err != nil {
			log.Errorw("Could not create deployment", "err", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		// Create Service
		err = resources.CreateServices(req.Name, req.Namespace, mysqlPort, appCtx)
		if err != nil {
			log.Errorw("Could not create deployment", "err", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Create response
		resp := types.CreateMySQLResponse {
			Message: "OK",
			Status: http.StatusOK,
		}

		if err := helper.WriteJSONResponse(w, resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}