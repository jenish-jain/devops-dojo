// I AM NOT DONE
package main

import (
	"fmt"
	"os"

	"github.com/jenish-jain/devops-dojo/internal/k8sexercise"
)

func main() {
	fmt.Println("Exercise: Debug a Pod Stuck in Pending State")
	fmt.Println("\n🔍 Debugging Scenario:")
	fmt.Println("A pod is stuck in 'Pending' state and won't schedule. Your task is to:")
	fmt.Println("  1. Identify why the scheduler can't place the pod")
	fmt.Println("  2. Fix the resource requests to reasonable values")
	fmt.Println("")

	runner := k8sexercise.NewRunner(
		"exercises/18_debug_pending/pending-pod.yaml",
		"Pod",
		&k8sexercise.PendingPodValidator{},
	)

	if err := runner.Run(); err != nil {
		os.Exit(1)
	}

	// Additional educational content
	fmt.Println("\n📚 What you learned:")
	fmt.Println("  • Pending state means the scheduler can't find a node for the pod")
	fmt.Println("  • Common causes: insufficient resources, node selectors, taints/tolerations")
	fmt.Println("  • Resource requests must be less than available node capacity")
	fmt.Println("  • Check Events section in 'kubectl describe pod' for scheduling failures")
	fmt.Println("  • Scheduler tries to place pods on nodes with enough resources")
	fmt.Println("\n🔧 Debugging commands:")
	fmt.Println("  kubectl get pods  # Check pod status")
	fmt.Println("  kubectl describe pod <pod-name>  # See scheduling events and errors")
	fmt.Println("  kubectl get nodes  # Check available nodes")
	fmt.Println("  kubectl describe nodes  # See node capacity and allocatable resources")
	fmt.Println("  kubectl top nodes  # See current resource usage (requires metrics-server)")
}
