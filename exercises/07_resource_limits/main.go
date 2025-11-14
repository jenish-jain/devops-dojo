// I AM NOT DONE
package main

import (
	"fmt"
	"os"

	"github.com/jenish-jain/devops-dojo/internal/k8sexercise"
)

func main() {
	fmt.Println("Exercise: Deploy a Pod with Resource Limits and Requests")

	runner := k8sexercise.NewRunner(
		"exercises/07_resource_limits/pod-resources.yaml",
		"Pod",
		&k8sexercise.ResourceLimitsValidator{},
	)

	if err := runner.Run(); err != nil {
		os.Exit(1)
	}

	// Additional educational content
	fmt.Println("\nWhat you learned:")
	fmt.Println("  • Resource requests: Minimum resources guaranteed for the pod")
	fmt.Println("  • Resource limits: Maximum resources the pod can use")
	fmt.Println("  • CPU is measured in millicores (m): 100m = 0.1 CPU core")
	fmt.Println("  • Memory is measured in bytes: Mi (mebibytes), Gi (gibibytes)")
	fmt.Println("  • Proper resource management prevents resource starvation")
	fmt.Println("\nTry these commands:")
	fmt.Println("  kubectl top pods  # View resource usage")
	fmt.Println("  kubectl describe pod <pod-name>  # See resource configuration")
}
