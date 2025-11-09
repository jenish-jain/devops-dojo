// I AM NOT DONE
package main

import (
	"fmt"
	"os"

	"github.com/jenish-jain/devops-dojo/internal/k8sexercise"
)

func main() {
	fmt.Println("Exercise: Create a HorizontalPodAutoscaler for automatic scaling")

	runner := k8sexercise.NewRunner(
		"exercises/15_hpa/hpa.yaml",
		"HorizontalPodAutoscaler",
		&k8sexercise.HPAValidator{},
	)

	if err := runner.Run(); err != nil {
		os.Exit(1)
	}

	// Additional educational content
	fmt.Println("\nWhat you learned:")
	fmt.Println("  • HPA automatically scales pods based on metrics")
	fmt.Println("  • Scales based on CPU, memory, or custom metrics")
	fmt.Println("  • Adjusts replica count between min and max")
	fmt.Println("  • Requires metrics-server to be installed")
	fmt.Println("  • Helps maintain performance during traffic spikes")
	fmt.Println("  • Reduces costs by scaling down during low traffic")
	fmt.Println("\nTry these commands:")
	fmt.Println("  kubectl get hpa")
	fmt.Println("  kubectl describe hpa <hpa-name>")
	fmt.Println("  kubectl top pods  # View current resource usage")
	fmt.Println("  # Generate load to test: kubectl run -i --tty load-generator --rm --image=busybox --restart=Never -- /bin/sh")
}
