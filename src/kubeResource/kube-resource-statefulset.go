package kubeResource

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type statefulSet struct {
	k8sClient *kubernetes.Clientset
}

func (stset *statefulSet) delete(namespace string, name string) (bool, error) {
	clientset := stset.k8sClient
	statefulsetClient := clientset.AppsV1beta1().StatefulSets(namespace)
	deletePolicy := metav1.DeletePropagationForeground
	err := statefulsetClient.Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy})
	if err != nil {

		return false, err
	}
	return true, err
}
