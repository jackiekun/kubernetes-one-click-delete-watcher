package kubeResource

type kubeResourceInterface interface {
	delete(namespace string, name string) (bool, error)
}
