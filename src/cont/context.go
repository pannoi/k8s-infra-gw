package cont

import (
	conf "infra-gw/src/configuration"
	"infra-gw/src/k8s"

	"go.uber.org/zap"
)

// Keys for context values
const (
	KeyReqID = iota
)

// AppContext is the application context holding references to things
type AppContext struct {
	Log    *zap.SugaredLogger
	K8s    *k8s.Client
	Config *conf.Config
}
