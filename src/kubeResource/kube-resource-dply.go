package kubeResource

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type deployment struct {
	k8sClient *kubernetes.Clientset
}

func (dply *deployment) delete(namespace string, name string) (bool, error) {
	clientset := dply.k8sClient
	deploymentsClient := clientset.ExtensionsV1beta1().Deployments(namespace)
	deletePolicy := metav1.DeletePropagationForeground
	err := deploymentsClient.Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy})
	if err != nil {
		return false, err
	}
	return true, nil
}
