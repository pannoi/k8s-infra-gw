package main

import (
	"context"
	"flag"
	"os"
	"time"

	"go.uber.org/zap"

	// local packages
	"infra-gw/src/api"
	conf "infra-gw/src/configuration"
	c "infra-gw/src/cont"
	"infra-gw/src/k8s"
	"infra-gw/src/util"
)

const version = "0.0.1"

var (
	config    conf.Config
	startTime int64
)

func main() {
	flag.StringVar(&config.ListenAddress, "listenAddress", "0.0.0.0:4444", "Address to bind web server to")
	flag.StringVar(&config.KubeConfig, "kubeconfig", "", "Path to kubeconfig (~/.kube/config if empty)")
	flag.BoolVar(&config.InCluster, "inCluster", false, "Use in-cluster config instead of kube config")
	flag.BoolVar(&config.DevMode, "devMode", false, "Enable DevMode")

	flag.Parse()

	startTime = time.Now().Unix()

	if config.KubeConfig == "" {
		config.KubeConfig = util.HomeDir() + "/.kube/config"
	}

	var logger *zap.Logger

	if config.DevMode {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}

	sugar := logger.Sugar()
	defer logger.Sync()

	sugar.With(
		"dev", config.DevMode,
		"kubeconfig", config.KubeConfig).Info("starting")

	rootCtx := context.Background()
	appCtx := &c.AppContext{
		Log:    sugar,
		K8s:    k8s.NewClientOrDie(config.KubeConfig, config.InCluster),
		Config: &config,
	}

	k8sVersion, err := appCtx.K8s.Clientset.ServerVersion()
	if err != nil {
		sugar.Errorf("Could not fetch k8s cluster version: %s", err)
	} else {
		sugar.Debugf("k8s version: %s", k8sVersion.String())
	}

	server := api.NewWeb(rootCtx, config.ListenAddress)
	server.SetupRoutes(rootCtx, appCtx)

	doneCh := server.Start(rootCtx, appCtx)
	sugar.Info("server started")

	err = <-doneCh
	if err != nil {
		sugar.Errorf("server stopped with err: %s", err)
	}
}

// ReadEnvOrPanic is a helper that returns an env var or panics if not set
func ReadEnvOrPanic(name string) string {
	val := os.Getenv(name)
	if val == "" {
		panic("Env var " + name + " must be set")
	}
	return val
}
