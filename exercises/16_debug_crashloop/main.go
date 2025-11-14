// I AM NOT DONE
package main

import (
	"fmt"
	"os"

	"github.com/jenish-jain/devops-dojo/internal/k8sexercise"
)

func main() {
	fmt.Println("Exercise: Debug a Pod in CrashLoopBackOff")
	fmt.Println("\n🔍 Debugging Scenario:")
	fmt.Println("A pod keeps crashing and restarting. Your task is to:")
	fmt.Println("  1. Identify why the pod is crashing")
	fmt.Println("  2. Fix the issue in the YAML file")
	fmt.Println("")

	runner := k8sexercise.NewRunner(
		"exercises/16_debug_crashloop/crashloop-pod.yaml",
		"Pod",
		&k8sexercise.CrashLoopValidator{},
	)

	if err := runner.Run(); err != nil {
		os.Exit(1)
	}

	// Additional educational content
	fmt.Println("\n📚 What you learned:")
	fmt.Println("  • CrashLoopBackOff means the pod crashed and Kubernetes is backing off restarts")
	fmt.Println("  • Use 'kubectl logs <pod-name>' to see why it crashed")
	fmt.Println("  • Use 'kubectl logs <pod-name> --previous' to see logs from crashed container")
	fmt.Println("  • Use 'kubectl describe pod <pod-name>' to see events")
	fmt.Println("  • Common causes: wrong command, missing files, application errors")
	fmt.Println("\n🔧 Debugging commands:")
	fmt.Println("  kubectl get pods  # Check pod status")
	fmt.Println("  kubectl describe pod <pod-name>  # See events and details")
	fmt.Println("  kubectl logs <pod-name>  # View logs")
	fmt.Println("  kubectl logs <pod-name> --previous  # View previous container logs")
}
