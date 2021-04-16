package configuration

// Config in k8s-gw config struct
type Config struct {
	ListenAddress string
	KubeConfig    string
	InCluster     bool
	DevMode       bool
}
