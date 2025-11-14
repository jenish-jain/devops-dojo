// I AM NOT DONE
package main

import (
	"fmt"
	"os"

	"github.com/jenish-jain/devops-dojo/internal/k8sexercise"
)

func main() {
	fmt.Println("Exercise: Add Liveness and Readiness Probes to a Pod")

	runner := k8sexercise.NewRunner(
		"exercises/08_health_probes/pod-probes.yaml",
		"Pod",
		&k8sexercise.HealthProbesValidator{},
	)

	if err := runner.Run(); err != nil {
		os.Exit(1)
	}

	// Additional educational content
	fmt.Println("\nWhat you learned:")
	fmt.Println("  • Liveness Probe: Kubernetes restarts the container if it fails")
	fmt.Println("  • Readiness Probe: Kubernetes stops sending traffic if it fails")
	fmt.Println("  • Probes can use httpGet, tcpSocket, or exec commands")
	fmt.Println("  • initialDelaySeconds: Wait before first probe")
	fmt.Println("  • periodSeconds: How often to probe")
	fmt.Println("  • Production apps should always have health probes!")
	fmt.Println("\nTry these commands:")
	fmt.Println("  kubectl describe pod <pod-name>  # See probe configuration")
	fmt.Println("  kubectl get events  # Watch probe failures")
	fmt.Println("  kubectl logs <pod-name>  # Check application logs")
}
