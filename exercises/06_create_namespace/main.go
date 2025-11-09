// I AM NOT DONE
package main

import (
	"fmt"
	"os"

	"github.com/jenish-jain/devops-dojo/internal/k8sexercise"
)

func main() {
	fmt.Println("Exercise: Create a Namespace for resource isolation")

	runner := k8sexercise.NewRunner(
		"exercises/06_create_namespace/namespace.yaml",
		"Namespace",
		&k8sexercise.NoOpValidator{},
	)

	if err := runner.Run(); err != nil {
		os.Exit(1)
	}

	// Additional educational content
	fmt.Println("\nWhat you learned:")
	fmt.Println("  • Namespaces provide logical isolation between resources")
	fmt.Println("  • They're useful for organizing resources by team, environment, or project")
	fmt.Println("  • Some resources are namespace-scoped (Pods, Services), others are cluster-scoped (Nodes)")
	fmt.Println("\nTry these commands:")
	fmt.Println("  kubectl get namespaces")
	fmt.Println("  kubectl get pods -n <namespace-name>")
	fmt.Println("  kubectl config set-context --current --namespace=<namespace-name>")
}
