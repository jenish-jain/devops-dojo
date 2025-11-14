// I AM NOT DONE
package main

import (
	"fmt"
	"os"

	"github.com/jenish-jain/devops-dojo/internal/k8sexercise"
)

func main() {
	fmt.Println("Exercise: Create a StatefulSet for stateful applications")

	runner := k8sexercise.NewRunner(
		"exercises/13_statefulset/statefulset.yaml",
		"StatefulSet",
		&k8sexercise.StatefulSetValidator{},
	)

	if err := runner.Run(); err != nil {
		os.Exit(1)
	}

	// Additional educational content
	fmt.Println("\nWhat you learned:")
	fmt.Println("  • StatefulSets manage stateful applications (databases, etc.)")
	fmt.Println("  • Pods have stable, unique network identities")
	fmt.Println("  • Ordered, graceful deployment and scaling")
	fmt.Println("  • Stable persistent storage via volumeClaimTemplates")
	fmt.Println("  • Pods named sequentially: <name>-0, <name>-1, <name>-2...")
	fmt.Println("  • Use for: databases, message queues, distributed systems")
	fmt.Println("\nTry these commands:")
	fmt.Println("  kubectl get statefulsets")
	fmt.Println("  kubectl get pods -l app=<your-app>")
	fmt.Println("  kubectl scale statefulset <name> --replicas=5")
}
