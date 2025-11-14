// I AM NOT DONE
package main

import (
	"fmt"
	"os"

	"github.com/jenish-jain/devops-dojo/internal/k8sexercise"
)

func main() {
	fmt.Println("Exercise: Create an Ingress for external HTTP/HTTPS access")

	runner := k8sexercise.NewRunner(
		"exercises/10_ingress/ingress.yaml",
		"Ingress",
		&k8sexercise.IngressValidator{},
	)

	if err := runner.Run(); err != nil {
		os.Exit(1)
	}

	// Additional educational content
	fmt.Println("\nWhat you learned:")
	fmt.Println("  • Ingress manages external HTTP/HTTPS access to services")
	fmt.Println("  • Provides load balancing, SSL termination, and name-based routing")
	fmt.Println("  • Requires an Ingress Controller (nginx, traefik, etc.)")
	fmt.Println("  • Routes traffic based on host and path rules")
	fmt.Println("  • More powerful than simple LoadBalancer services")
	fmt.Println("\nTry these commands:")
	fmt.Println("  kubectl get ingress")
	fmt.Println("  kubectl describe ingress <ingress-name>")
	fmt.Println("  curl -H 'Host: your-app.example.com' <ingress-ip>")
}
