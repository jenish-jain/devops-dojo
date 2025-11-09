// I AM NOT DONE
package main

import (
	"fmt"
	"os"

	"github.com/jenish-jain/devops-dojo/internal/k8sexercise"
)

func main() {
	fmt.Println("Exercise: Create and use a ConfigMap for application configuration")

	runner := k8sexercise.NewRunner(
		"exercises/04_create_configmap/configmap.yaml",
		"ConfigMap",
		&k8sexercise.ConfigMapValidator{},
	)

	if err := runner.Run(); err != nil {
		os.Exit(1)
	}

	// Additional educational content for successful ConfigMap creation
	fmt.Println("\nTry these commands to explore your ConfigMap:")
	fmt.Println("  kubectl get configmaps")
	fmt.Println("  kubectl describe configmap/<configmap-name>")
	fmt.Println("  kubectl get configmap/<configmap-name> -o yaml")
}
