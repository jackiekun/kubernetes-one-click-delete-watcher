package kubeResource

import (
	"fmt"
	"log"
	"time"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

const (
	deletingResourceNameLabel = "mappingResourceName"
)

func getClientSet() *kubernetes.Clientset {
	// out of cluster setting

	// var kubeconfig *string
	// if home := homedir.HomeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }
	// flag.Parse()

	// config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// clientset, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// incluster setting
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// end

	return clientset
}

func factory(name string) kubeResourceInterface {
	switch name {
	case "deployment":
		return &deployment{k8sClient: getClientSet()}
	case "statefulset":
		return &statefulSet{k8sClient: getClientSet()}
	default:
		panic("Such a resouce is not defined")
	}
}

func Run(namespace string, resourceType string) {

	var deletingResouce kubeResourceInterface
	deletingResouce = factory(resourceType)

	client := getClientSet()
	watchlist := cache.NewListWatchFromClient(client.Core().RESTClient(), "services", namespace, fields.Everything())
	_, controller := cache.NewInformer(
		watchlist,
		&v1.Service{},
		time.Second*0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {

			},
			DeleteFunc: func(obj interface{}) {
				service := obj.(*v1.Service)
				log.Printf(fmt.Sprintf("Get service %s deletion event", service.Name))
				deletingResouceName := service.Labels[deletingResourceNameLabel]
				_, err := deletingResouce.delete(namespace, deletingResouceName)
				if err != nil {
					log.Printf("Failed to delete resource %s, the error is %s\n", deletingResouceName, err.Error())
				} else {
					log.Printf("Deleted resource %s\n", deletingResouceName)
				}
				// 	pm.deleteDeployment(DEFAULTNAMESPACE, serviceName)
				// pm.deleteStatefulset(DEFAULTNAMESPACE, serviceName)

			},
			UpdateFunc: func(oldObj, newObj interface{}) {
			},
		},
	)

	stop := make(chan struct{})
	go controller.Run(stop)
	for {
		time.Sleep(time.Second)
	}
}
