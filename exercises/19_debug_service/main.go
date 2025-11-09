// I AM NOT DONE
package main

import (
	"fmt"
	"os"

	"github.com/jenish-jain/devops-dojo/internal/k8sexercise"
)

func main() {
	fmt.Println("Exercise: Debug Service Not Routing Traffic")
	fmt.Println("\n🔍 Debugging Scenario:")
	fmt.Println("A service exists but traffic isn't reaching the pods. Your task is to:")
	fmt.Println("  1. Identify why the service can't find the pods")
	fmt.Println("  2. Fix the service selector to match pod labels")
	fmt.Println("")
	fmt.Println("NOTE: This exercise requires both a Service and a Pod/Deployment.")
	fmt.Println("Create the service.yaml with correct selectors that match your pods.")
	fmt.Println("")

	runner := k8sexercise.NewRunner(
		"exercises/19_debug_service/service.yaml",
		"Service",
		&k8sexercise.ServiceDebugValidator{},
	)

	if err := runner.Run(); err != nil {
		os.Exit(1)
	}

	// Additional educational content
	fmt.Println("\n📚 What you learned:")
	fmt.Println("  • Services route traffic to pods using label selectors")
	fmt.Println("  • Service selector must EXACTLY match pod labels")
	fmt.Println("  • Check 'kubectl get endpoints' to see if service found any pods")
	fmt.Println("  • Empty endpoints means selector doesn't match any pods")
	fmt.Println("  • Port and targetPort must match the container's listening port")
	fmt.Println("\n🔧 Debugging commands:")
	fmt.Println("  kubectl get svc  # List services")
	fmt.Println("  kubectl describe svc <service-name>  # See selector and endpoints")
	fmt.Println("  kubectl get endpoints <service-name>  # See pod IPs service routes to")
	fmt.Println("  kubectl get pods --show-labels  # See all pod labels")
	fmt.Println("  kubectl get pods -l app=myapp  # Test selector query")
	fmt.Println("\n💡 Pro tip: If endpoints are empty, the selector doesn't match!")
}
