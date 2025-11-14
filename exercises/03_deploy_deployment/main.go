// I AM NOT DONE
package main

import (
	"fmt"
	"os"

	"github.com/jenish-jain/devops-dojo/internal/k8sexercise"
)

func main() {
	fmt.Println("Exercise: Deploy a Kubernetes Deployment with replicas")

	runner := k8sexercise.NewRunner(
		"exercises/03_deploy_deployment/deployment.yaml",
		"Deployment",
		&k8sexercise.DeploymentValidator{},
	)

	if err := runner.Run(); err != nil {
		os.Exit(1)
	}

	// Additional educational content for successful deployment
	fmt.Println("\nTry these commands to see your deployment in action:")
	fmt.Println("  kubectl get deployments")
	fmt.Println("  kubectl get replicasets")
	fmt.Println("  kubectl get pods")
	fmt.Println("  kubectl scale deployment/<deployment-name> --replicas=5")
}
