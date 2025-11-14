// I AM NOT DONE
package main

import (
	"fmt"
	"os"

	"github.com/jenish-jain/devops-dojo/internal/k8sexercise"
)

func main() {
	fmt.Println("Exercise: Create a Secret for sensitive data")

	runner := k8sexercise.NewRunner(
		"exercises/05_create_secret/secret.yaml",
		"Secret",
		&k8sexercise.SecretValidator{},
	)

	if err := runner.Run(); err != nil {
		os.Exit(1)
	}

	// Additional educational content for successful Secret creation
	fmt.Println("\nTry these commands to explore your Secret:")
	fmt.Println("  kubectl get secrets")
	fmt.Println("  kubectl describe secret/<secret-name>")
	fmt.Println("  kubectl get secret/<secret-name> -o yaml")
	fmt.Println("\n⚠️  Remember: Secrets in git are NOT secret! Use proper secret management.")
}
