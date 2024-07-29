package clientsvc

import (
	"os"

	"canarails.dev/services/envsvc"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func initConfig() *rest.Config {
	configPath := os.Getenv(envsvc.K8S_CONFIG_PATH)

	if len(configPath) == 0 {
		// 未配置 config path 时使用 in cluster 连接
		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err)
		}

		return config
	} else {
		// 通过 config path 创建 client 连接
		config, err := clientcmd.BuildConfigFromFlags("", configPath)
		if err != nil {
			panic(err)
		}

		return config
	}
}

func New() *kubernetes.Clientset {
	clientset, err := kubernetes.NewForConfig(initConfig())
	if err != nil {
		panic(err)
	}

	return clientset
}

func NewDynamic() *dynamic.DynamicClient {
	clientset, err := dynamic.NewForConfig(initConfig())
	if err != nil {
		panic(err)
	}

	return clientset
}
