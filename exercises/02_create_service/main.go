// I AM NOT DONE
package main

import (
	"fmt"
	"os"

	"github.com/jenish-jain/devops-dojo/internal/k8sexercise"
)

func main() {
	fmt.Println("Creating a Kubernetes service using kubectl")

	runner := k8sexercise.NewRunner(
		"exercises/02_create_service/service.yaml",
		"Service",
		&k8sexercise.NoOpValidator{},
	)

	if err := runner.Run(); err != nil {
		os.Exit(1)
	}
}
