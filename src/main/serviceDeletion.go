package main

import (
	"../kubeResource"
)

const (
	Namespace                  = "default"
	CascadeDeletionResouceType = "deployment"
)

func main() {
	kubeResource.Run(Namespace, CascadeDeletionResouceType)
}
